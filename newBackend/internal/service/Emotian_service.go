package service

import (
	"net/http"

	"forum_backend/internal/project_error"
	"forum_backend/internal/repository"
	"forum_backend/model"
)

type IEmotianService interface {
	CreateOrUpdateEmotionComment(model.Like) error
	CreateOrUpdateEmotionPost(model.Like) error
}

type EmotianService struct {
	repo repository.IEmotionSQL
}

func NewEmotianService(repoEmotian repository.IEmotionSQL) IEmotianService {
	return &EmotianService{
		repoEmotian,
	}
}

func (e *EmotianService) CreateOrUpdateEmotionPost(postEmo model.Like) error {
	if postEmo.Islike < -1 && postEmo.Islike > 1 {
		return project_error.NewUserError("Invalid like", http.StatusBadRequest)
	}
	exists, err := e.repo.EmotionPostExists(postEmo.PostID, postEmo.UserID)
	if err != nil {
		return project_error.NewInternalError("Failed to check if emotion exists")
	}
	if exists {
		err = e.repo.UpdateEmotionPost(postEmo)
		if err != nil {
			return project_error.NewInternalError("Failed to update emotion")
		}
	} else {
		err = e.repo.CreateEmotionPost(postEmo)
		if err != nil {
			return project_error.NewInternalError("Failed to create emotion")
		}
	}

	return nil
}

func (e *EmotianService) CreateOrUpdateEmotionComment(commentEmo model.Like) error {
	if commentEmo.Islike < -1 && commentEmo.Islike > 1 {
		return project_error.NewUserError("Invalid like", http.StatusBadRequest)
	}
	exists, err := e.repo.EmotionCommentExists(commentEmo.CommentID, commentEmo.UserID)
	if err != nil {
		return project_error.NewUserError("Invalid like", http.StatusBadRequest)
	}
	if exists {
		err = e.repo.UpdateEmotionComment(commentEmo)
		if err != nil {
			return project_error.NewInternalError("Failed to update emotion")
		}
	} else {
		err = e.repo.CreateEmotionComment(commentEmo)
		if err != nil {
			return project_error.NewInternalError("Failed to create emotion")
		}
	}

	return nil
}
