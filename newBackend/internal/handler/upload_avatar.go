package handler

import (
	"io/ioutil"
	"net/http"
)

func (h *Handler) uploadAvatar(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(int)
	file, header, err := r.FormFile("image")
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO magic value
	if header.Size > 5*1024*1024 {
		http.Error(w, "Image size exceeds the limit of 10 MB", http.StatusBadRequest)
		return
	}
	defer file.Close()
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.services.IAuth_service.UploadAvatar(imageData, id); err != nil {
		h.Logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Avatar uploaded successfully"))
}
