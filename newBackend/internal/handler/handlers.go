package handler

import (
	"net/http"

	"forum_backend/internal/Log"
	"forum_backend/internal/service"
	"forum_backend/model"
)

type Handler struct {
	services *service.Service
	Router   *http.ServeMux
	Logger   *Log.Logger
}
type Response struct {
	Metadata model.MetaDataPost `json:"metadata"`
	Posts    []model.Post       `json:"posts"`
}

func NewHandler(s *service.Service, logger *Log.Logger) *Handler {
	return &Handler{
		services: s,
		Router:   http.NewServeMux(),
		Logger:   logger,
	}
}

func (h Handler) Start() http.Handler {
	h.Router.HandleFunc("/api", h.home)                                          // ok
	h.Router.HandleFunc("/api/auth/sign_in", h.signIn)                           // ok
	h.Router.HandleFunc("/api/auth/check_token", h.checkToken)                   // ok
	h.Router.HandleFunc("/api/auth/sign_up", h.signUp)                           // ok
	h.Router.HandleFunc("/api/auth/log_out", h.logOut)                           // ok
	h.Router.HandleFunc("/api/auth/upload_avatar", h.CheckCredo(h.uploadAvatar)) // ok
	h.Router.HandleFunc("/api/post/get_post/", h.getPost)                        // ok
	h.Router.HandleFunc("/api/post/get_my_posts/", h.CheckCredo(h.getMyPost))
	h.Router.HandleFunc("/api/post/get_my_liked_posts/", h.CheckCredo(h.getMyLikedPost))
	h.Router.HandleFunc("/api/post/get_post_category", h.getPostCategory)
	h.Router.HandleFunc("/api/post/create_post", h.CheckCredo(h.createPost))
	h.Router.HandleFunc("/api/post/get_comment/", h.getComment)
	h.Router.HandleFunc("/api/post/create_comment", h.CheckCredo(h.createComment))
	h.Router.HandleFunc("/api/emotian/comment", h.CheckCredo(h.emotianComment))
	h.Router.HandleFunc("/api/emotian/post", h.CheckCredo(h.emotianPost))
	h.Router.HandleFunc("/api/personal_info", h.CheckCredo(h.personalInfo))

	return h.CorsHeaders(h.Router)
}
