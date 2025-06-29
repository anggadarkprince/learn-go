package orm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// https://gorm.io/docs/query.html
func TestQuerySingleObject(t *testing.T) {
	user := User{}
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	err := db.First(&user).Error
	assert.Nil(t, err, "Expected no error when querying single object")

	user = User{}
	// SELECT * FROM `users` ORDER BY `users`.`id` DESC LIMIT 1
	err = db.Last(&user).Error
	assert.Nil(t, err, "Expected no error when querying last object")
}

func TestQueryNotFound(t *testing.T) {
	user := User{}
	// SELECT * FROM `users` WHERE id = 9999 ORDER BY `users`.`id` LIMIT 1
	err := db.First(&user, 9999).Error // should return record not found error
	if errors.Is(err, gorm.ErrRecordNotFound) {
  		fmt.Printf("Not Found: %v", err)
	}
	assert.NotNil(t, err, "Expected error when querying non-existent object")
	assert.Equal(t, "record not found", err.Error(), "Expected record not found error")
}

func TestQueryInlineCondition(t *testing.T) {
	user := User{}
	// SELECT * FROM `users` WHERE id = 1 ORDER BY `users`.`id` LIMIT 1
	err := db.First(&user, "id = ?", 1).Error // with order by
	assert.Nil(t, err, "Expected no error when querying with inline condition")

	user = User{}
	// SELECT * FROM `users` WHERE id = 1 LIMIT 1
	err = db.Take(&user, "id = ?", 1).Error // without order by
	assert.Nil(t, err, "Expected no error when querying with inline condition using Take")
}

func TestQueryGetAll(t *testing.T) {
	var users []User
	// SELECT * FROM `users`
	err := db.Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying all objects")

	// SELECT * FROM `users` WHERE id in (1,2,3)
	err = db.Find(&users, "id in ?", []int{1, 2, 3}).Error
	//err = db.Where("id > ?", 0).Find(&users).Error // also works
	assert.Nil(t, err, "Expected no error when querying with condition")
	assert.Greater(t, len(users), 0, "Expected to find some users")
}

func TestQueryCondition(t *testing.T) {
	var users []User
	// SELECT * FROM `users` WHERE id > 2 AND status = 'PENDING'
	err := db.Where("id > ?", 2).Where("status = ?", "PENDING").Find(&users).Error // where should be chained before Find
	assert.Nil(t, err, "Expected no error when querying with condition")
	assert.Greater(t, len(users), 0, "Expected to find some users")

	// Test with multiple conditions
	// SELECT * FROM `users` WHERE id > 2 AND name LIKE '%Keenan%'
	err = db.Where("id > ? AND name LIKE ?", 2, "%Keenan%").Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with multiple conditions")
	assert.Greater(t, len(users), 0, "Expected to find some users matching multiple conditions")
}

func TestOrOperator(t *testing.T) {
	var users []User
	// SELECT * FROM `users` WHERE id > 2 OR status = 'ACTIVATED'
	err := db.Where("id > ?", 2).Or("status = ?", "ACTIVATED").Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with OR operator")
	assert.Greater(t, len(users), 0, "Expected to find some users matching OR condition")
}

func TestNotOperator(t *testing.T) {
	var users []User
	err := db.Not("id = ?", 1).Where("status = ?", "PENDING").Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with NOT operator")
	assert.Greater(t, len(users), 0, "Expected to find some users not matching condition")

	// Test with multiple NOT conditions
	// SELECT * FROM `users` WHERE NOT (id > 1 OR status = 'ACTIVATED')
	err = db.Not("id > ? OR status = ?", 1, "ACTIVATED").Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with multiple NOT conditions")
	assert.Greater(t, len(users), 0, "Expected to find some users not matching multiple conditions")
}

func TestGroupingCondition(t *testing.T) {
	var users []User
	// SELECT * FROM `users` WHERE id > 2 OR (status = 'PENDING' AND name LIKE '%Keenan%')
	err := db.Where("id > ?", 2).Or(db.Where("status = ?", "PENDING").Where("name LIKE ?", "%Keenan%")).Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with grouped conditions")
	assert.Greater(t, len(users), 0, "Expected to find some users matching grouped conditions")
}

func TestSelectFields(t *testing.T) {
	var users []User
	// SELECT id, name FROM `users`
	err := db.Select("id, name").Find(&users).Error
	assert.Nil(t, err, "Expected no error when selecting specific fields")
	assert.Greater(t, len(users), 0, "Expected to find some users with selected fields")

	for _, user := range users {
		assert.NotEmpty(t, user.ID, "Expected user ID to be not empty")
		assert.NotEmpty(t, user.Name, "Expected user Name to be not empty")
		assert.Empty(t, user.Username, "Expected user Username to be empty when selecting specific fields")
		assert.Empty(t, user.Email, "Expected user Email to be empty when selecting specific fields")
	}

	// Test with conditions
	// SELECT id, name FROM `users` WHERE status = 'PENDING'
	err = db.Select("id, name").Where("status = ?", "PENDING").Find(&users).Error
	assert.Nil(t, err, "Expected no error when selecting specific fields with condition")
	assert.Greater(t, len(users), 0, "Expected to find some users with selected fields and condition")
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name:   "Angga",
		Status: "PENDING",
		Email:  "", // Not included in the query, because it's default value of string, use map condition if you want to include it
	}
	var users []User
	// SELECT * FROM `users` WHERE status = 'PENDING' AND name LIKE '%Angga%'
	err := db.Where(userCondition).Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with struct condition")
	assert.Greater(t, len(users), 0, "Expected to find some users matching struct condition")

	// Test with multiple struct conditions
	// SELECT * FROM `users` WHERE `users`.`status` = 'PENDING' OR `users`.`name` = 'Keenan'
	err = db.Where(User{Status: "PENDING"}).Or(User{Name: "Keenan"}).Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with multiple struct conditions")
	assert.Greater(t, len(users), 0, "Expected to find some users matching multiple struct conditions")
}

func TestMapCondition(t *testing.T) {
	// SELECT * FROM `users` WHERE status = 'PENDING' AND name LIKE '%Angga%'
	condition := map[string]any{
		"username": "", // This will be included in the query, even if it's empty
	}
	var users []User
	// SELECT * FROM `users` WHERE `users`.`username` = ''
	err := db.Where(condition).Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with map condition")
	assert.Greater(t, len(users), 0, "Expected to find some users matching map condition")

	// Test with multiple map conditions
	// SELECT * FROM `users` WHERE `users`.`status` = 'PENDING' OR `users`.`name` = 'Keenan'
	err = db.Where(map[string]any{"status": "PENDING"}).Or(map[string]interface{}{"name": "Keenan"}).Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with multiple map conditions")
	assert.Greater(t, len(users), 0, "Expected to find some users matching multiple map conditions")
}

func TestQueryWithLimitAndOffset(t *testing.T) {
	var users []User
	// SELECT * FROM `users` ORDER BY username asc, email desc LIMIT 2 OFFSET 1
	err := db.Order("username asc, email desc").Limit(2).Offset(1).Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with limit and offset")
	assert.Equal(t, 2, len(users), "Expected to find exactly 2 users with limit and offset")

	// Test with only limit
	// SELECT * FROM `users` LIMIT 3
	err = db.Limit(3).Find(&users).Error
	assert.Nil(t, err, "Expected no error when querying with only limit")
	assert.GreaterOrEqual(t, len(users), 3, "Expected to find at least 3 users with limit")
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func TestQueryNonModel(t *testing.T) {
	var count int64
	// SELECT COUNT(*) FROM `users`
	err := db.Model(&User{}).Count(&count).Error
	assert.Nil(t, err, "Expected no error when counting users")
	assert.Greater(t, count, int64(0), "Expected to find some users")

	var userResponses []UserResponse
	// Test with custom query
	// SELECT `users`.`id`,`users`.`name`,`users`.`username` FROM `users` WHERE status = 'PENDING'
	err = db.Model(&User{}).Where("status = ?", "PENDING").Find(&userResponses).Error
	assert.Nil(t, err, "Expected no error when counting users with condition")
	assert.Greater(t, count, int64(0), "Expected to find some users with status PENDING")
}

// https://gorm.io/docs/update.html#Updates-multiple-columns
func TestUpdateUsingSave(t *testing.T) {
	user := User{}
	// SELECT * FROM `users` WHERE `users`.`id` = 1 LIMIT 1
	err := db.Take(&user, 1).Error
	assert.Nil(t, err, "Expected no error when querying user for update")

	// Update user status
	user.Name = "Angga"
	user.Status = "ACTIVATED"
	// UPDATE `users` SET `name`='Angga',`username`='angga.ari',`email`='angga@mail.com',`password`='secret',`status`='ACTIVATED',`updated_at`='2025-06-29 14:23:15.723' WHERE `id` = 1
	err = db.Save(&user).Error // Update all fields including auto-updated fields
	assert.Nil(t, err, "Expected no error when updating user status")

	// Verify update
	updatedUser := User{}
	err = db.Take(&updatedUser, 1).Error
	assert.Nil(t, err, "Expected no error when querying updated user")
	assert.Equal(t, "ACTIVATED", updatedUser.Status, "Expected user status to be updated to ACTIVATED")
}

func TestUpdateSelectedColumn(t *testing.T) {
	// Update single column from empty model
	// UPDATE `users` SET `email`='angga@mail.com' WHERE id = 1
	err := db.Model(&User{}).Where("id = ?", 1).Update("email", "angga@mail.com").Error
	assert.Nil(t, err, "Expected no error when updating user email")

	// UPDATE `users` SET `name`='Angga',`status`='ACTIVATED' WHERE id = 1
	err = db.Model(&User{}).Where("id = ?", 1).Updates(map[string]any{
		"name":   "Angga",
		"status": "ACTIVATED",
	}).Error
	assert.Nil(t, err, "Expected no error when updating user name and status")

	// Update from existing model
	user := User{}
	// SELECT * FROM `users` WHERE `users`.`id` = 1 LIMIT 1
	err = db.Take(&user, 1).Error
	assert.Nil(t, err, "Expected no error when querying user for update")

	// UPDATE `users` SET `name`='Angga',`status`='ACTIVATED',`email`='' WHERE `id` = 1
	err = db.Model(&user).Select("Name", "Email", "Status").Updates(User{
		Name:   "Angga",
		Status: "ACTIVATED",
		// Email will set to default value of string because it's mentioned in Select
	}).Error
	assert.Nil(t, err, "Expected no error when updating selected columns")

	// UPDATE `users` SET `email`='angga@mail.com' WHERE `id` = 1
	err = db.Model(&user).Updates(User{
		Name:   "Angga",
		Status: "", // Will be ignored because it's passing default value of string, use map or update specific fields if you want to update it
	}).Error
	assert.Nil(t, err, "Expected no error when updating selected columns")

	// Update only specific fields
	// UPDATE `users` SET `email`='angga@mail.com' WHERE `id` = 1
	err = db.Model(&user).Update("email", "angga@mail.com").Error
	assert.Nil(t, err, "Expected no error when updating user status")

	// Verify update
	updatedUser := User{}
	err = db.Take(&updatedUser, 1).Error
	assert.Nil(t, err, "Expected no error when querying updated user")
	assert.Equal(t, "ACTIVATED", updatedUser.Status, "Expected user status to be updated to ACTIVATED")
}
