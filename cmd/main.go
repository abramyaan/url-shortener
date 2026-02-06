package main

import (
	"fmt"
	"net/http"
	"url-shortener/configs"
	"url-shortener/internal/auth"
	"url-shortener/internal/user"
	"url-shortener/pkg/db"
	"url-shortener/pkg/jwt"
)

func main() {
	cfg := configs.LoadConfig()
	database := db.NewDb(cfg)
	database.AutoMigrate(&user.User{})
	userRepo := user.NewUserRepository(database)
	jwtTool := jwt.NewJWT(cfg.Auth.Secret)
	authService := auth.NewAuthService(userRepo, jwtTool)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      cfg,
		AuthService: authService,
	})
	fmt.Printf("🚀 Сервер запущен на порту %s\n", cfg.Port)
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %v\n", err)
	}
}
