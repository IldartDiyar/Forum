package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) getPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var userID int
	if tokenClient, err := r.Cookie("token"); err == nil { // TODO error handle?
		user, _ := h.services.GetUserByToken(tokenClient.Value)
		userID = user.ID
	}

	postID, err := strconv.Atoi(r.URL.Query().Get("id"))
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

	if metadatePost.Posts < postID || postID <= 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	post, err := h.services.IPost_service.GetPost(postID, userID)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
