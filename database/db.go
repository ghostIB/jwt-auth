package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"jwt-auth/models"
)

// DB - глобальна змінна для підключення до бази даних
var DB *gorm.DB

func Connect() {
	// рядок підключення до PostgreSQL
	dsn := "host=localhost user=jwt_user password=1234 dbname=jwt_auth port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не вдалось підключитись до БД: ", err)
	}

	fmt.Println("Підключення до БД успішне!")

	// AutoMigrate автоматично створить таблицю users якщо її немає
	db.AutoMigrate(&models.User{})

	DB = db
}
