package orm

import (
	"orm/models"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

func TestCreateUser(t *testing.T) {
	user := models.User{
		ID: 1,
		Name: "Angga",
		Username: "angga.ari",
		Email: "angga@mail.com",
		Password: "secret",
		Information: "this is additional info",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []models.User
	for i := 2; i <= 10; i++ {
		users = append(users, models.User{
			ID: uint64(i),
			Name: "User " + strconv.Itoa(i),
			Username: "user_" + strconv.Itoa(i),
			Email: "user_" + strconv.Itoa(i) + "@mail.com",
			Password: "secret",
		})
	}

	result := db.Create(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(9), result.RowsAffected)

	//db.CreateInBatches(&users, 5) // batch insert each 5 data (chunking)
}

func TestUpdateUser(t *testing.T) {
	user := models.User{
		ID: 1,
		Name: "Angga Ari",
		Username: "angga.ari",
		Email: "", // Not included when passing default type value, use map instead
	}
	response := db.Model(&user).Where("id = ?", user.ID).Updates(models.User{Name: user.Name, Email: user.Email})
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected, "Expected one row to be updated")
}

func TestAutoIncrement(t *testing.T) {
	for i := 1; i <= 10; i++ {
		userLog := models.UserLog{
			UserID:    i,
			Action:    "test action",
		}

		err := db.Create(&userLog).Error
		assert.Nil(t, err, "Expected no error when creating user log")
		assert.NotZero(t, userLog.ID, "Expected auto-incremented ID to be")
	}
}

func TestSaveOrUpdate(t *testing.T) {
	userLog := models.UserLog{
		UserID: 1,
		Action: "Data Updated",
	}
	err := db.Save(&userLog).Error // Insert
	assert.Nil(t, err, "Expected no error when saving or updating user")


	userLog.UserID = 2
	err = db.Save(&userLog).Error // Update
	assert.Nil(t, err, "Expected no error when saving or updating user")
}

func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
	userLog := models.UserLog{
		ID: 1, // Assuming this ID already exists
		UserID: 1,
		Action: "Data Updated",
	}
	err := db.Save(&userLog).Error // Update
	assert.Nil(t, err, "Expected no error when saving or updating user log")

	userLog.ID = 99 // Reset ID to zero for insert
	// Try to update first, if it doesn't exist, it will insert
	// UPDATE `user_logs` SET `user_id`=1,`action`='Data Updated',`created_at`=0,`updated_at`=1751184667246 WHERE `id` = 99
	// Affected rows: 0
	// INSERT INTO `user_logs` (`user_id`,`action`,`created_at`,`updated_at`,`id`) VALUES (1,'Data Updated',1751184667246,1751184667246,99) ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`),`action`=VALUES(`action`)
	err = db.Save(&userLog).Error // Insert
	assert.Nil(t, err, "Expected no error when saving or updating user log")
}

func TestConflict(t *testing.T) {
	user := models.User{
		ID: 88,
		Name: "Angga Conflict",
	}

	// https://gorm.io/gen/create.html#Upsert-On-Conflict
	// This will insert a new user or update the existing one if there's a conflict on the primary key
	// If the user with ID 88 already exists, it will update the Name field
	err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&user).Error
	assert.Nil(t, err, "Expected no error when handling conflict")
}


func TestDeleteUser(t *testing.T) {
	// Delete by loading the record first
	var user models.User
	db.Take(&user, 9) // Assuming this user exists
	assert.NotZero(t, user.ID, "Expected user with ID 9 to exist")

	err := db.Delete(&user).Error
	assert.Nil(t, err, "Expected no error when deleting user")

	// Delete by ID without loading the record first
	userId := models.User{
		ID: 10, // Assuming this user exists
	}
	response := db.Delete(&userId)
	//response := db.Delete(&userId, "id = ?", userId.ID)
	//response := db.Where("id = ?", userId.ID).Delete(&User{})
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected, "Expected one row to be deleted")

	// You can also delete by ID directly
	// response = db.Delete(&User{}, 1)
	// assert.Nil(t, response.Error)
	// assert.Equal(t, int64(1), response.RowsAffected, "Expected one row to be deleted")
}

func TestSoftDelete(t *testing.T) {
	// Soft delete by loading the record first
	todo := models.Todo{
		UserID: 1,
		Title: "Test Todo",
		Description: "This is a test todo",
	}
	err := db.Create(&todo).Error
	assert.Nil(t, err, "Expected no error when creating todo")

	// UPDATE `todos` SET `deleted_at`='2025-06-29 16:15:13.831' WHERE `todos`.`id` = 1 AND `todos`.`deleted_at` IS NULL
	db.Delete(&todo) // Soft delete
	assert.NotZero(t, todo.ID, "Expected todo ID to be set after soft delete")
	assert.NotNil(t, todo.DeletedAt, "Expected DeletedAt to be set after soft delete")

	var todos []models.Todo
	// SELECT * FROM `todos` WHERE `todos`.`deleted_at` IS NULL
	err = db.Find(&todos).Error
	assert.Nil(t, err, "Expected no error when querying todos after soft delete")
	assert.Equal(t, 0, len(todos), "Expected no todos to be returned after soft delete")
}

func TestUnscoped(t *testing.T) {
	// Unscoped delete to permanently remove the record
	todo := models.Todo{
		UserID: 1,
		Title: "Test Todo",
		Description: "This is a test todo",
	}
	err := db.Create(&todo).Error
	assert.Nil(t, err, "Expected no error when creating todo")

	// Permanently delete the todo
	db.Unscoped().Delete(&todo)
	assert.NotZero(t, todo.ID, "Expected todo ID to be set after unscoped delete")

	// Get all todos, including soft-deleted ones
	var todos []models.Todo
	err = db.Unscoped().Find(&todos).Error
	assert.Nil(t, err, "Expected no error when querying todos after unscoped delete")
}