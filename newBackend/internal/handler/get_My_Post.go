package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"forum_backend/internal/project_error"
)

func (h *Handler) getMyPost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(int)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusUnsupportedMediaType)
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
	metadatePost, err := h.services.IPost_service.GetMetaDataPost(category)
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
	if metadatePost.Pages < page || page <= 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	metadatePost, err = h.services.IPost_service.GetMetaDataMyPost(id, category)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	posts, err := h.services.IPost_service.GetMyPosts(id, page, category)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(Response{metadatePost, posts})
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
