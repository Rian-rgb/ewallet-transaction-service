package helper

import (
	"ewallet-transaction/internal/domain/transaction"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func SetupPostgreSQL() {
	var err error

	host := GetEnv("DB_HOST", "")
	user := GetEnv("DB_USER", "")
	pass := GetEnv("DB_PASSWORD", "")
	name := GetEnv("DB_NAME", "")
	port := GetEnv("DB_PORT", "")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, pass, name, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database ", err)
	}

	logrus.Info("Successfully connect to database")

	err = DB.AutoMigrate(&transaction.Entity{})
	if err != nil {
		return
	}
}
