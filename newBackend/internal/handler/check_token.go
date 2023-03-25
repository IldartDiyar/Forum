package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) checkToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	tokenClient, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		h.Logger.Error(err.Error())
		return
	}
	user, err := h.services.IAuth_service.GetUserByToken(tokenClient.Value)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		h.Logger.Error(err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		h.Logger.Error(err.Error())
		return
	}
}
