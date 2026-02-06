package auth

import (
	"encoding/json"
	"net/http"
	"url-shortener/configs"
	"url-shortener/internal/user"
	"url-shortener/pkg/res"
)

// 1. Структура самого Хендлера
type AuthHandler struct {
	Config      *configs.Config
	AuthService *AuthService
}

// 2. Структура для передачи зависимостей (Deps)
type AuthHandlerDeps struct {
	Config      *configs.Config
	AuthService *AuthService
}

// 3. Конструктор и регистратор маршрутов
func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

// 4. Структура DTO (Data Transfer Object)
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string     `json:"token"`
	User  *user.User `json:"user"`
}

// 5. Метод регистрации
func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			res.Json(w, "invalid request", http.StatusBadRequest)
			return
		}

		user, err := h.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Json(w, user, http.StatusCreated)
	}
}

// 6. Метод логина (заглушка)
func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body LoginRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			res.Json(w, "Неккоретный формат данных", http.StatusBadRequest)
			return
		}
		user, token, err := h.AuthService.Login(body.Email, body.Password)
		if err != nil {
			res.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}

		res.Json(w, LoginResponse{
			Token: token,
			User:  user,
		}, http.StatusOK)
	}
}
