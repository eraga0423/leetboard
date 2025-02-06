package adapters

import (
	"encoding/json"
	"net/http"

	"1337b0rd/internal/core"
)

type PostHandler struct {
	service *core.PostService
}

func NewPostHandler(service *core.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Content string `json:"content"`
		Image   string `json:"image"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	post, err := h.service.CreatePost(req.Content, req.Image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(post)
}
