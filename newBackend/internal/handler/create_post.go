package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"forum_backend/internal/project_error"
	"forum_backend/model"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
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
	var post model.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.services.IPost_service.InsertPost(post, id); err != nil {
		switch err.(type) {
		case *project_error.UserError:
			userErr := err.(*project_error.UserError)
			http.Error(w, userErr.Error(), userErr.Status())
		case *project_error.InternalError:
			internalErr := err.(*project_error.InternalError)
			h.Logger.Error(internalErr.Error())
			http.Error(w, internalErr.Error(), http.StatusInternalServerError)
		default:
			serverErr := project_error.NewServerError("Internal Server Error")
			h.Logger.Error(serverErr.Error())
			http.Error(w, serverErr.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}
