package middleware

import (
	"context"
	"net/http"
	"strings"
	"url-shortener/pkg/jwt"
	"url-shortener/pkg/res"
)

type contextKey string

const UserIdKey contextKey = "userId"

func Auth(next http.Handler, jwtService *jwt.JWT) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			res.Json(w, "Нужен токен", http.StatusUnauthorized)
			return 
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			res.Json(w, "Неверный формат токена", http.StatusUnauthorized)
			return 
		}
		token:= tokenParts[1]
		claims, err := jwtService.Parse(token)
		if err != nil {
			res.Json(w,"Токен не валиден", http.StatusUnauthorized)
			return 
		}
		userId, ok := claims["userId"].(float64)
		if !ok {
			res.Json(w, "Ошибка данных в токене", http.StatusUnauthorized)
			return 
		}
		ctx:=context.WithValue(r.Context(), UserIdKey, uint(userId))
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}