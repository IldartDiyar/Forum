package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) getComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var userID int
	tokenClient, err := r.Cookie("token")
	if err == nil {
		user, _ := h.services.GetUserByToken(tokenClient.Value)
		userID = user.ID
	}

	commentID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	metadatePost, err := h.services.IPost_service.GetMetaDataPost("")
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if metadatePost.Posts < commentID || commentID <= 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	comment, err := h.services.IComment_service.GetComments(commentID, userID)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
