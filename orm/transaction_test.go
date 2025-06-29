package orm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func (tx *gorm.DB) error {
		err := tx.Create(&User{
			Name: "Keenan",
			Username: "keenan",
			Email: "keenan@mail.com",
			Password: "secret",
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			Name: "Evander",
			Username: "evander",
			Email: "evander@mail.com",
			Password: "secret",
		}).Error
		if err != nil {
			return err
		}

		return nil
	})
	assert.Nil(t, err)
}

func TestTransactionRollback(t *testing.T) {
	err := db.Transaction(func (tx *gorm.DB) error {
		err := tx.Create(&User{
			Name: "Alastair",
			Username: "alastair",
			Email: "alastair@mail.com",
			Password: "secret",
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			Name: "Evander",
			Username: "evander", // trigger error unique column
			Email: "evander@mail.com",
			Password: "secret",
		}).Error
		if err != nil {
			return err
		}

		return nil
	})
	assert.NotNil(t, err)
}

func TestManualTransaction(t *testing.T) {
	tx := db.Begin()
	defer tx. Rollback() // even it's success, the rollback it's not affected

	err1 := tx.Create(&User{
		Name: "Khusina",
		Username: "khusina",
		Email: "khusina@mail.com",
		Password: "secret",
	}).Error // make sure this is rollback

	err2 := tx.Create(&User{
		Name: "Evander",
		Username: "evander", // trigger error unique column
		Email: "evander@mail.com",
		Password: "secret",
	}).Error

	if err1 == nil && err2 == nil {
		tx.Commit()
	}
}

func TestLock(t *testing.T) {
	tx := db.Transaction(func (tx *gorm.DB) error {
		var user User
		// SELECT * FROM `users` WHERE `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1 FOR UPDATE
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, 1).Error
		if err != nil {
			return err
		}

		// PDATE `users` SET `name`='Angga Ari',`username`='angga.ari',`email`='angga@mail.com',`password`='secret',`status`='ACTIVATED',`updated_at`='2025-06-29 16:32:55.658' WHERE `id` = 1
		user.Name = "Angga Ari"
		user.Email = "angga@mail.com"
		err = tx.Save(&user).Error
		return err
	})
	assert.Nil(t, tx, "Expected no error when manually handling transaction")
}