package handler

import (
	"context"
	"forum_backend/model"
	"net/http"
	"time"
)

func (h *Handler) CorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cookie, Content-Length, Accept-Encoding, X-CSRF-Token, charset, Credentials, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			// w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckCredo(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenClient, err := r.Cookie("token")
		if err != nil {
			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var user model.User
		if user, err = h.services.GetUserByToken(tokenClient.Value); err != nil {
			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user.Token = tokenClient.Value
		if !user.TokenDuration.Before(time.Now()) {
			if err := h.services.DeleteToken(user.Token); err != nil {
				h.Logger.Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userID", user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
