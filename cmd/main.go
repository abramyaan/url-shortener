package main

import (
	"net/http"
	"url-shortener/configs"
	"url-shortener/internal/auth"
	"url-shortener/internal/event" // Добавили
	"url-shortener/internal/link"
	"url-shortener/internal/stat"
	"url-shortener/internal/user"
	"url-shortener/pkg/db"
	"url-shortener/pkg/jwt"
)

func main() {
	cfg := configs.LoadConfig()
	database := db.NewDb(cfg)


	database.AutoMigrate(&user.User{}, &link.Link{}, &event.Event{})

	jwtTool := jwt.NewJWT(cfg.Auth.Secret)
	router := http.NewServeMux()

	userRepo := user.NewUserRepository(database)
	authService := auth.NewAuthService(userRepo, jwtTool)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      cfg,
		AuthService: authService,
	})


	eventRepo := event.NewEventRepository(database.DB)

	
	linkRepo := link.NewLinkRepository(database.DB)
	linkService := link.NewLinkService(linkRepo, eventRepo)
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkService: linkService,
		JwtService:  jwtTool,
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
    LinkRepo:   linkRepo,
    JwtService: jwtTool,
})


	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	println("Server started on :8080")
	server.ListenAndServe()
}