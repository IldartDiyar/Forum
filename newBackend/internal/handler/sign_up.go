package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"forum_backend/internal/project_error"
	"forum_backend/model"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Invalid or empty authorization data", http.StatusBadRequest)
		return
	}

	var usr model.User
	if err = json.Unmarshal(body, &usr); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err = h.services.CreateUser(usr); err != nil {
		h.Logger.Error(err.Error())

		switch e := err.(type) {
		case *project_error.UserError:
			http.Error(w, e.Error(), e.Status())
		case *project_error.InternalError:
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
		default:
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}
