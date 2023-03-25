package service

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"forum_backend/internal/project_error"
	"forum_backend/internal/repository"
	"forum_backend/model"
)

type IPost_service interface {
	GetPost(int, int) (model.Post, error)
	InsertPost(model.Post, int) error
	GetMetaDataPost(string) (model.MetaDataPost, error)
	GetMetaDataMyLikedPost(int, string) (model.MetaDataPost, error)
	GetMetaDataMyPost(int, string) (model.MetaDataPost, error)
	GetAllPost(int, string) ([]model.Post, error)
	GetCategory() ([]model.Category, error)
	GetMyPosts(int, int, string) ([]model.Post, error)
	GetMyLikedPosts(int, int, string) ([]model.Post, error)
}

type Post_service struct {
	repo repository.IPostSQL
}

func NewPost_service(repoPost repository.IPostSQL) IPost_service {
	return &Post_service{
		repoPost,
	}
}

func (p *Post_service) GetPost(postId int, userId int) (model.Post, error) {
	return p.repo.GetPost(postId, userId)
}

func (p *Post_service) GetMyPosts(userId int, page int, category string) ([]model.Post, error) {
	var categoryId int
	if len(category) != 0 {
		var err error
		categoryId, err = p.repo.GetCategoryID(category)
		if err != nil {
			return []model.Post{}, err
		}
	}
	pageSize := 10
	offset := (page - 1) * pageSize

	return p.repo.GetMyPost(userId, offset, categoryId)
}

func (p *Post_service) GetMetaDataMyLikedPost(id int, category string) (model.MetaDataPost, error) {
	var categoryId int
	if len(category) != 0 {
		var err error
		categoryId, err = p.repo.GetCategoryID(category)
		if errors.Is(err, sql.ErrNoRows) {
			return model.MetaDataPost{}, project_error.NewUserError("wrong category", http.StatusNotFound)
		}
		if err != nil {
			return model.MetaDataPost{}, err
		}
	}
	return p.repo.GetMetaDataMyLikedPost(id, categoryId)
}

func (p *Post_service) GetMetaDataPost(category string) (model.MetaDataPost, error) {
	var categoryId int
	if len(category) != 0 {
		var err error
		categoryId, err = p.repo.GetCategoryID(category)
		if errors.Is(err, sql.ErrNoRows) {
			return model.MetaDataPost{}, project_error.NewUserError("wrong category", http.StatusNotFound)
		}
		if err != nil {
			return model.MetaDataPost{}, err
		}
	}
	return p.repo.GetMetaDataPost(categoryId)
}

func (p *Post_service) GetMetaDataMyPost(id int, category string) (model.MetaDataPost, error) {
	var categoryId int
	if len(category) != 0 {
		var err error
		categoryId, err = p.repo.GetCategoryID(category)
		if errors.Is(err, sql.ErrNoRows) {
			return model.MetaDataPost{}, project_error.NewUserError("wrong category", http.StatusNotFound)
		}
		if err != nil {
			return model.MetaDataPost{}, err
		}
	}
	return p.repo.GetMetaDataMyPost(id, categoryId)
}

func (p *Post_service) GetAllPost(page int, category string) ([]model.Post, error) {
	var categoryId int
	if len(category) != 0 {
		var err error
		categoryId, err = p.repo.GetCategoryID(category)
		if err != nil {
			return []model.Post{}, err
		}
	}
	pageSize := 10
	offset := (page - 1) * pageSize
	return p.repo.GetAllPost(offset, categoryId)
}

func (p *Post_service) InsertPost(post model.Post, id int) error {
	username, err := p.repo.GetUsernameByID(id)
	if err != nil {
		return project_error.NewInternalError("Error getting username: " + err.Error())
	}
	if username != post.Author {
		return project_error.NewUserError("Author not same Username", http.StatusBadRequest)
	}
	if len(strings.Trim(post.Title, " ")) == 0 && len(strings.Trim(post.Content, " ")) == 0 {
		return project_error.NewUserError("Empty Title or Body", http.StatusBadRequest)
	}
	CategoryId, err := p.repo.GetCategoryID(post.Category)
	if err != nil {
		return project_error.NewInternalError("Error getting category ID: " + err.Error())
	}
	post.CategoryId = CategoryId
	if len(post.ImageData) != 0 {
		imageBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(post.ImageData))
		if err != nil {
			return project_error.NewUserError("Invalid image data", http.StatusBadRequest)
		}

		post.ImageData, err = resize(imageBytes, 500, 500)
		if err != nil {
			return project_error.NewInternalError("Error resizing image: " + err.Error())
		}
	}
	if err := p.repo.CreatePost(post); err != nil {
		return project_error.NewInternalError("Error creating post: " + err.Error())
	}
	return nil
}

func (p *Post_service) GetCategory() ([]model.Category, error) {
	return p.repo.GetCategory()
}

func (p *Post_service) GetMyLikedPosts(userId int, page int, category string) ([]model.Post, error) {
	var categoryId int
	if len(category) != 0 {
		var err error
		categoryId, err = p.repo.GetCategoryID(category)
		if err != nil {
			return []model.Post{}, err
		}
	}
	pageSize := 10
	offset := (page - 1) * pageSize
	return p.repo.GetMyLikedPost(userId, offset, categoryId)
}
