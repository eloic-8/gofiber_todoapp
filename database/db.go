package database

import (
	"fmt"
	"os"

	"github.com/eloicahhing/go-fiber-tutorial/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB - global DB var
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {

	var err error
	// dsn TODO: comfirm at the end is "?charset=utf8mb4&parseTime=True&loc=Local" necessary
	dsn := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	if err != nil {
		panic(err)
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&models.Todo{})
	DB.AutoMigrate(&models.Project{})
	//DB.AutoMigrate(&models.LoyaltyCard{})
	//DB.AutoMigrate(&models.LoyaltyCardWallet{})

	fmt.Println("Database Migrated")
}
