package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"forum_backend/internal/project_error"
	"forum_backend/model"
)

func (h *Handler) emotianPost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(int)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		h.Logger.Error("Couldn't read the body of a request in SignInHandler or body is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var like model.Like
	like.UserID = id
	err = json.Unmarshal(body, &like)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.services.IEmotianService.CreateOrUpdateEmotionPost(like); err != nil {
		h.Logger.Error(err.Error())
		if userErr, ok := err.(*project_error.UserError); ok {
			http.Error(w, userErr.Error(), userErr.Status())
			return
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
