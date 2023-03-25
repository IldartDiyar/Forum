package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum_backend/internal/project_error"
	"forum_backend/model"
)

type homeResponse struct {
	Metadata model.MetaDataPost `json:"metadata"`
	Posts    []model.Post       `json:"posts"`
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	page := 1
	if r.URL.Query().Has("page") {
		pageStr, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			http.NotFound(w, r)
		}
		page = pageStr
	}

	category := r.URL.Query().Get("category")
	metadata, err := h.services.IPost_service.GetMetaDataPost(category)
	if userErr, ok := err.(*project_error.UserError); ok {
		h.Logger.Error(err.Error())
		http.Error(w, userErr.Error(), userErr.Status())
		return
	}
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if metadata.Pages < page || page <= 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	posts, err := h.services.IPost_service.GetAllPost(page, category)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := homeResponse{
		Metadata: metadata,
		Posts:    posts,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
