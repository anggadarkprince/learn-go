package orm

import (
	"fmt"
	"orm/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestCreateWallet(t *testing.T) {
	// Create a new wallet
	wallet := &models.Wallet{
		UserID: 1,
		Balance: 100.00,
	}

	if err := db.Create(wallet).Error; err != nil {
		t.Fatalf("Failed to create wallet: %v", err)
	}

	if wallet.ID == 0 {
		t.Fatal("Wallet ID should not be zero after creation")
	}
}

func TestRetrieveRelation(t *testing.T) {
	// Retrieve a user and their wallet
	var user models.User
	// SELECT * FROM `wallets` WHERE `wallets`.`user_id` = 1
	// SELECT * FROM `users` WHERE `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1
	if err := db.Preload("Wallet").First(&user, 1).Error; err != nil {
		t.Fatalf("Failed to retrieve user with wallet: %v", err)
	}

	if user.Wallet.ID == 0 {
		t.Fatal("User's wallet should not have zero ID")
	}
	if user.Wallet.UserID != uint64(user.ID) {
		t.Fatal("User's wallet UserID should match user's ID")
	}
}

func TestRetrieveRelationJoin(t *testing.T) {
	// Retrieve users with their wallets using join
	var users []models.User
	// SELECT `users`.`id`,`users`.`name`,`users`.`username`,`users`.`email`,`users`.`password`,`users`.`status`,`users`.`created_at`,`users`.`updated_at`,`Wallet`.`id` AS `Wallet__id`,`Wallet`.`user_id` AS `Wallet__user_id`,`Wallet`.`balance` AS `Wallet__balance`,`Wallet`.`created_at` AS `Wallet__created_at`,`Wallet`.`updated_at` AS `Wallet__updated_at` FROM `users` LEFT JOIN `wallets` `Wallet` ON `users`.`id` = `Wallet`.`user_id` WHERE users.id = '1' LIMIT 1
	if err := db.Joins("Wallet").Take(&users, "users.id = ?", "1").Error; err != nil {
		t.Fatalf("Failed to retrieve users with wallets: %v", err)
	}

	for _, user := range users {
		if user.Wallet.ID == 0 {
			t.Fatal("User's wallet should not have zero ID")
		}
		if user.Wallet.UserID != uint64(user.ID) {
			t.Fatal("User's wallet UserID should match user's ID")
		}
	}
}

func TestAutoCreateUpdate(t *testing.T) {
	// Create a new user
	user := models.User{
		Name: "Auto User",
		Username: "autouser",
		Email: "user@mail.com",
		Password: "password",
		Wallet: models.Wallet{ // Create a wallet with auto timestamps or update
			// If the user already exists, it will update the wallet balance
			Balance: 50.00,
		},
	}
	// INSERT INTO `users` (`name`,`username`,`email`,`password`,`status`,`created_at`,`updated_at`) VALUES ('Auto User','autouser','user@mail.com','password','PENDING','2025-06-30 21:04:21.188','2025-06-30 21:04:21.188')
	// INSERT INTO `wallets` (`user_id`,`balance`,`created_at`,`updated_at`) VALUES (100,50,'2025-06-30 21:04:21.197','2025-06-30 21:04:21.197') ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`)
	err := db.Create(&user).Error;
	if err != nil {
		t.Fatalf("Failed to create user with auto timestamps: %v", err)
	}
}

func TestSkipAutoCreateUpdate(t *testing.T) {
	// Create a new user without auto timestamps
	user := models.User{
		Name: "Skip Auto User",
		Username: "skipautouser",
		Email: "userskip@mail.com",
		Password: "password",
		Wallet: models.Wallet{ // Ignore wallet relation
			Balance: 50,
		},
	}
	// INSERT INTO `users` (`name`,`username`,`email`,`password`,`status`,`created_at`,`updated_at`) VALUES ('Skip Auto User','skipautouser','userskip@mail.com','password','PENDING','2025-06-30 21:07:51.978','2025-06-30 21:07:51.978')
	err := db.Omit(clause.Associations).Create(&user).Error;
	if err != nil {
		t.Fatalf("Failed to create user with auto timestamps: %v", err)
	}
}

func TestUserAndAddresses(t *testing.T) {
	// Create a new user
	user := models.User{
		Name: "Auto User Address",
		Username: "autouseraddress",
		Email: "autouseraddress@mail.com",
		Password: "password",
		Addresses: []models.Address{
			{Address: "Jl semarang 10"},
			{Address: "Jl sumatra 23"},
		},
	}
	// INSERT INTO `users` (`name`,`username`,`email`,`password`,`status`,`created_at`,`updated_at`) VALUES ('Auto User Address','autouseraddress','autouseraddress@mail.com','password','PENDING','2025-06-30 21:15:32.217','2025-06-30 21:15:32.217')
	// NSERT INTO `addresses` (`user_id`,`address`,`created_at`,`updated_at`) VALUES (102,'Jl semarang 10','2025-06-30 21:15:32.222','2025-06-30 21:15:32.222'),(102,'Jl sumatra 23','2025-06-30 21:15:32.222','2025-06-30 21:15:32.222') ON DUPLICATE KEY UPDATE `user_id`=VALUES(`user_id`)
	err := db.Create(&user).Error;
	if err != nil {
		t.Fatalf("Failed to create user with auto timestamps: %v", err)
	}
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var usersPreload []models.User
	// SELECT * FROM `addresses` WHERE `addresses`.`user_id` IN (1,2,3,4,5,6,7,8,11,12,18,100,101,102)
	// ELECT `users`.`id`,`users`.`name`,`users`.`username`,`users`.`email`,`users`.`password`,`users`.`status`,`users`.`created_at`,`users`.`updated_at`,`Wallet`.`id` AS `Wallet__id`,`Wallet`.`user_id` AS `Wallet__user_id`,`Wallet`.`balance` AS `Wallet__balance`,`Wallet`.`created_at` AS `Wallet__created_at`,`Wallet`.`updated_at` AS `Wallet__updated_at` FROM `users` LEFT JOIN `wallets` `Wallet` ON `users`.`id` = `Wallet`.`user_id`
	err := db.Preload("Addresses").Joins("Wallet").Find(&usersPreload).Error
	assert.Nil(t, err)

	var user models.User;
	err = db.Preload("Addresses").Joins("Wallet").Take(&user, 1).Error
	assert.Nil(t, err)
}

func TestBelongsTo(t *testing.T) {
	fmt.Println("Preload")
	var addresses []models.Address
	// SELECT * FROM `addresses`
	// SELECT * FROM `users` WHERE `users`.`id` IN (`101`, `102`)
	err := db.Preload("User").Find(&addresses).Error
	assert.Nil(t, err)

	fmt.Println("Join")
	addresses = []models.Address{} // Reset
	// SELECT `addresses`.`id`,`addresses`.`user_id`,`addresses`.`address`,`addresses`.`created_at`,`addresses`.`updated_at`,`User`.`id` AS `User__id`,`User`.`name` AS `User__name`,`User`.`username` AS `User__username`,`User`.`email` AS `User__email`,`User`.`password` AS `User__password`,`User`.`status` AS `User__status`,`User`.`created_at` AS `User__created_at`,`User`.`updated_at` AS `User__updated_at` FROM `addresses` LEFT JOIN `users` `User` ON `addresses`.`user_id` = `User`.`id`
	err = db.Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
}

func TestBelongsToWallet(t *testing.T) {
	fmt.Println("Preload")
	var wallets []models.Wallet
	// SELECT * FROM `wallets`
	// SELECT * FROM `users` WHERE `users`.`id` IN (1,100)
	err := db.Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("Join")
	wallets = []models.Wallet{} // Reset
	// SELECT `wallets`.`id`,`wallets`.`user_id`,`wallets`.`balance`,`wallets`.`created_at`,`wallets`.`updated_at`,`User`.`id` AS `User__id`,`User`.`name` AS `User__name`,`User`.`username` AS `User__username`,`User`.`email` AS `User__email`,`User`.`password` AS `User__password`,`User`.`status` AS `User__status`,`User`.`created_at` AS `User__created_at`,`User`.`updated_at` AS `User__updated_at` FROM `wallets` LEFT JOIN `users` `User` ON `wallets`.`user_id` = `User`.`id`
	err = db.Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}

func TestCreateManyToMany(t *testing.T) {
	product := models.Product{
		ID: 1,
		Name: "Iphone 16 Pro",
		Price: 999,
	}
	// INSERT INTO `products` (`name`,`price`,`created_at`,`updated_at`,`id`) VALUES ('Iphone 16 Pro',999,'2025-06-30 21:52:15.851','2025-06-30 21:52:15.851',1)
	err := db.Create(&product).Error
	assert.Nil(t, err)

	// INSERT INTO `user_like_products` (`product_id`,`user_id`) VALUES ('1','1')
	err = db.Table("user_like_products").Create(map[string]any{
		"user_id": "1",
		"product_id": "1",
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToMany(t *testing.T) {
	var product models.Product
	// SELECT * FROM `users` WHERE `users`.`id` = 1
	// SELECT * FROM `user_like_products` WHERE `user_like_products`.`product_id` = 1
	// SELECT * FROM `products` WHERE `products`.`id` = 1 LIMIT 1
	err := db.Preload("LikedByUsers").Take(&product, 1).Error
	assert.Nil(t, err)
}

func TestPreloadManyToManyUser(t *testing.T) {
	var user models.User
	// SELECT * FROM `products` WHERE `products`.`id` = 1
	// SELECT * FROM `user_like_products` WHERE `user_like_products`.`user_id` = 1
	// SELECT * FROM `users` WHERE `users`.`id` = 1 LIMIT 1
	err := db.Preload("LikeProducts").Take(&user, 1).Error
	assert.Nil(t, err)
}

func TestAssociationFind(t *testing.T) {
	var product models.Product
	err := db.Take(&product, 1).Error
	assert.Nil(t, err)

	var users []models.User
	// SELECT `users`.`id`,`users`.`name`,`users`.`username`,`users`.`email`,`users`.`password`,`users`.`status`,`users`.`created_at`,`users`.`updated_at` FROM `users` JOIN `user_like_products` ON `user_like_products`.`user_id` = `users`.`id` AND `user_like_products`.`product_id` = 1 WHERE name LIKE 'A%'
	err = db.Model(&product).Where("username LIKE ?", "A%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
}

func TestAssociationAppend(t *testing.T) {
	// SELECT * FROM `users` WHERE `users`.`id` = 1 LIMIT 1
	var user models.User
	err := db.Take(&user, 1).Error
	assert.Nil(t, err)

	// SELECT * FROM `products` WHERE `products`.`id` = 1 LIMIT 1
	var product models.Product
	err = db.Take(&product, 1).Error
	assert.Nil(t, err)

	// INSERT INTO `user_like_products` (`product_id`,`user_id`) VALUES (1,1) ON DUPLICATE KEY UPDATE `product_id`=`product_id`
	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user models.User
		err := tx.First(&user, 1).Error
		assert.Nil(t, err)

		wallet := models.Wallet{
			UserID: user.ID,
			Balance: 10,
		}

		// Error, gorm will set null for old wallet if the user_id cannot be null
		return db.Model(&user).Association("Wallet").Replace(&wallet)
	})
	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	// SELECT * FROM `users` WHERE `users`.`id` = 1 LIMIT 1
	var user models.User
	err := db.Take(&user, 1).Error
	assert.Nil(t, err)

	// SELECT * FROM `products` WHERE `products`.`id` = 1 LIMIT 1
	var product models.Product
	err = db.Take(&product, 1).Error
	assert.Nil(t, err)

	// DELETE FROM `user_like_products` WHERE `user_like_products`.`product_id` = 1 AND `user_like_products`.`user_id` = 1
	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)
}

func TestAssociationClear(t *testing.T) {
	// SELECT * FROM `products` WHERE `products`.`id` = 1 LIMIT 1
	var product models.Product
	err := db.Take(&product, 1).Error
	assert.Nil(t, err)

	// Remove relation data
	// DELETE FROM `user_like_products` WHERE `user_like_products`.`product_id` = 1
	err = db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}