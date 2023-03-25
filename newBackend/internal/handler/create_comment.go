package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"forum_backend/model"
)

func (h *Handler) createComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.Context().Value("userID").(int)
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		h.Logger.Error("Couldn't read the body of a request in SignInHandler or body is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var comment model.Comment
	err = json.Unmarshal(body, &comment)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.services.IComment_service.InsertComment(comment, id)
	if err == errors.New("Empty Body") {
		w.WriteHeader(http.StatusBadRequest)
		h.Logger.Error(err.Error())
		return
	}
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
