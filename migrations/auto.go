package main

import (
	"os"
	"os/user"
	"url-shortener/internal/link"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Не удалось загрузить .env файл")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic("Не удалось подключиться к базе : " + err.Error())
	}
	err = db.AutoMigrate(&user.User{}, &link.Link{})
	if err != nil {
		panic("Ошибка миграции: " + err.Error())
	}
	println("Миграция успешно завершена !")
}
