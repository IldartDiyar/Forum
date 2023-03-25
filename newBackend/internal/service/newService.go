package service

import (
	"forum_backend/internal/repository"
)

type Service struct {
	IAuth_service
	IComment_service
	IEmotianService
	IPost_service
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		IAuth_service:    NewAuth_service(repo.IAuthSQL),
		IComment_service: NewComment_service(repo.ICommentSQL),
		IEmotianService:  NewEmotianService(repo.IEmotionSQL),
		IPost_service:    NewPost_service(repo.IPostSQL),
	}
}
