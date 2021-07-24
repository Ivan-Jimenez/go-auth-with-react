package database

import (
	"github.com/Ivan-Jimenez/go-auth-with-react/server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:root@/go_auth"), &gorm.Config{})
	if err != nil {
		panic("Failed database connection!")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
