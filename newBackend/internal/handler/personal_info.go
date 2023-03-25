package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) personalInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(int)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	Info, err := h.services.IAuth_service.PersonalInfo(id)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(Info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Logger.Error(err.Error())
		return
	}
}
