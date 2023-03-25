package service

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"forum_backend/internal/project_error"
	"forum_backend/internal/repository"
	"forum_backend/model"

	"github.com/gofrs/uuid"
)

type IAuth_service interface {
	CreateUser(model.User) error
	GetUserByToken(string) (model.User, error)
	CheckUser(model.User) (model.User, error)
	DeleteToken(string) error
	UploadAvatar([]byte, int) error
	PersonalInfo(int) (model.User, error)
}

type Auth_service struct {
	repo repository.IAuthSQL
}

func NewAuth_service(repoAuth repository.IAuthSQL) IAuth_service {
	return &Auth_service{
		repoAuth,
	}
}

func (a *Auth_service) CreateUser(user model.User) error {
	if err := isValidEmail(user.Email); err != nil {
		return project_error.NewUserError(err.Error(), http.StatusBadRequest)
	}
	if err := isValidUsername(user.Username); err != nil {
		return project_error.NewUserError(err.Error(), http.StatusBadRequest)
	}
	if err := isValidPassword(user.Password); err != nil {
		return project_error.NewUserError(err.Error(), http.StatusBadRequest)
	}
	_, err := a.repo.GetUserByUsername(user.Username)
	if err == nil {
		return project_error.NewUserError("username already exists", http.StatusConflict)
	} else if !errors.Is(err, sql.ErrNoRows) {
		return project_error.NewServerError(err.Error())
	}

	_, err = a.repo.GetUserByEmail(user.Email)
	if err == nil {
		return project_error.NewUserError("email already exists", http.StatusConflict)
	} else if !errors.Is(err, sql.ErrNoRows) {
		return project_error.NewServerError(err.Error())
	}

	if err := a.repo.CreateUser(user); err != nil {
		return project_error.NewServerError(err.Error())
	}
	return nil
}

func (a *Auth_service) GetUserByToken(token string) (model.User, error) {
	var user model.User
	user, err := a.repo.CheckUserByToken(token)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (a *Auth_service) CheckUser(user model.User) (model.User, error) {
	if err := isValidUsername(user.Username); err != nil {
		return user, project_error.NewUserError("invalid username", http.StatusBadRequest)
	}
	if err := isValidPassword(user.Password); err != nil {
		return user, project_error.NewUserError("invalid password", http.StatusBadRequest)
	}
	user, err := a.repo.CheckUser(user)
	if err != nil {
		return user, project_error.NewUserError("User does not exist", http.StatusUnauthorized)
	}
	token, err := uuid.NewV4()
	if err != nil {
		return user, project_error.NewInternalError("error generating token")
	}
	user.Token = token.String()
	user.TokenDuration = time.Now().Add(72 * time.Hour)
	a.repo.DeleteTokenById(user.ID)
	if err := a.repo.SaveToken(user); err != nil {
		return user, project_error.NewInternalError("error saving token to database")
	}
	return user, nil
}

func (a *Auth_service) DeleteToken(token string) error {
	if _, err := a.repo.CheckUserByToken(token); err != nil {
		return err
	}
	return a.repo.DeleteToken(token)
}

func (a *Auth_service) UploadAvatar(file []byte, UserId int) error {
	resized, err := resize(file, 100, 100)
	if err != nil {
		return err
	}
	if err = a.repo.UploadAvatar(resized, UserId); err != nil {
		return a.repo.UpdateAvatar(resized, UserId)
	}
	return nil
}

func (a *Auth_service) PersonalInfo(userId int) (model.User, error) {
	return a.repo.GetUserInfo(userId)
}
