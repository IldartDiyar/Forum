package service

import (
	"net/http"
	"strings"

	"forum_backend/internal/project_error"
	"forum_backend/internal/repository"
	"forum_backend/model"
)

type IComment_service interface {
	GetComments(int, int) ([]model.Comment, error)
	InsertComment(model.Comment, int) error
}

type Comment_service struct {
	repo repository.ICommentSQL
}

func NewComment_service(repoComment repository.ICommentSQL) IComment_service {
	return &Comment_service{
		repoComment,
	}
}

func (c *Comment_service) GetComments(postId int, userId int) ([]model.Comment, error) {
	return c.repo.GetComment(postId, userId)
}

func (c *Comment_service) InsertComment(comment model.Comment, id int) error {
	username, err := c.repo.GetUsernameById(id)
	if err != nil {
		return project_error.NewServerError(err.Error())
	}

	if username != comment.Author {
		return project_error.NewUserError("author does not match username", http.StatusBadRequest)
	}

	if len(strings.TrimSpace(comment.Body)) == 0 {
		return project_error.NewUserError("comment body cannot be empty", http.StatusBadRequest)
	}

	if err := c.repo.CreateComment(comment); err != nil {
		return project_error.NewServerError(err.Error())
	}

	return nil
}
