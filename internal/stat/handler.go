package stat

import (
	"net/http"
	"url-shortener/internal/link"
	"url-shortener/pkg/jwt"
	"url-shortener/pkg/middleware"
	"url-shortener/pkg/res"
)

type StatHandler struct {
	LinkRepo *link.LinkRepository
}

type StatHandlerDeps struct {
	LinkRepo   *link.LinkRepository
	JwtService *jwt.JWT
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		LinkRepo: deps.LinkRepo,
	}

	// Маршрут защищен, так как статистику может смотреть только владелец
	router.Handle("GET /stat", middleware.Auth(handler.GetStat(), deps.JwtService))
}



func (h *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(middleware.UserIdKey).(uint)
		if !ok {
			res.Json(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Ищем все ссылки этого пользователя
		// (Нам нужно добавить метод GetByUserID в link/repository.go)
		links, err := h.LinkRepo.GetByUserID(userID)
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var response []link.StatResponse
		for _, l := range links {
			response = append(response, link.StatResponse{
				LinkID: l.ID,
				Hash:   l.Hash,
				Clicks: len(l.Events), // GORM подтянет события через Preload
			})
		}

		res.Json(w, response, http.StatusOK)
	}
}