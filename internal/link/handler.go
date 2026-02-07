package link

import (
	"encoding/json"
	"net/http"
	"url-shortener/pkg/jwt"
	"url-shortener/pkg/middleware"
	"url-shortener/pkg/res"
)

type LinkHandler struct {
	LinkService *LinkService
}

type LinkHandlerDeps struct {
	LinkService *LinkService
	JwtService  *jwt.JWT
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkService: deps.LinkService,
	}

	router.Handle("POST /link", middleware.Auth(handler.Create(), deps.JwtService))
	router.HandleFunc("GET /go/{hash}", handler.GoTo())
}



func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body LinkCreateRequest

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			res.Json(w, "invalid request body", http.StatusBadRequest)
			return
		}

		userID, ok := r.Context().Value(middleware.UserIdKey).(uint)
		if !ok {
			res.Json(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		link, err := h.LinkService.Create(body.URL, userID)
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, LinkResponse{
			ID:   link.ID,
			URL:  link.Url,
			Hash: link.Hash,
		}, http.StatusCreated)
	}
}

func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if hash == "" {
			res.Json(w, "hash is required", http.StatusBadRequest)
			return
		}

		link, err := h.LinkService.GetByHash(hash)
		if err != nil {
			res.Json(w, "link not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}