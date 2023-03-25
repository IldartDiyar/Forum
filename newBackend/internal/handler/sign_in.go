package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"forum_backend/internal/project_error"
	"forum_backend/model"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		h.Logger.Error("Couldn't read the body of a request in SignInHandler or body is empty")
		http.Error(w, "Invalid or empty authorization data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var usr model.User
	if err = json.Unmarshal(body, &usr); err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, "Invalid authorization data", http.StatusBadRequest)
		return
	}

	usr, err = h.services.CheckUser(usr)
	if err != nil {
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

	cookie := &http.Cookie{
		Name:     "token",
		Value:    usr.Token,
		Expires:  usr.TokenDuration,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	if err = json.NewEncoder(w).Encode(usr); err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
}
