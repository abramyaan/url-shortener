package main

import (
	"fmt"
	"net/http"
	"url-shortener/configs"
	"url-shortener/internal/auth"
	"url-shortener/internal/user"
	"url-shortener/pkg/db"
)

func main() {
	cfg := configs.LoadConfig()
	database := db.NewDb(cfg)
	database.AutoMigrate(&user.User{})
	userRepo := user.NewUserRepository(database)
	authService := auth.NewAuthService(userRepo)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: cfg,
		AuthService: authService,
	})
	fmt.Printf("🚀 Сервер запущен на порту %s\n", cfg.Port)
	server := &http.Server {
		Addr: ":" + cfg.Port,
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %v\n" , err)
	}
}