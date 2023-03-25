package handler

import "net/http"

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	tokenClient, err := r.Cookie("token")
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err := h.services.IAuth_service.DeleteToken(tokenClient.Value); err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "null",
		MaxAge:   -1,
		HttpOnly: false,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
}
