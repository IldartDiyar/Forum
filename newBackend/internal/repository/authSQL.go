package repository

import (
	"database/sql"
	"time"

	"forum_backend/model"
)

type AuthSQL struct {
	db *sql.DB
}

type IAuthSQL interface {
	CreateUser(model.User) error
	CheckUser(model.User) (model.User, error)
	CheckUserByToken(string) (model.User, error)
	SaveToken(model.User) error
	DeleteToken(string) error
	UploadAvatar(string, int) error
	GetAvatar(int) (string, error)
	UpdateAvatar(string, int) error
	DeleteTokenById(int)
	GetUserInfo(int) (model.User, error)
	GetUserByEmail(string) (model.User, error)
	GetUserByUsername(string) (model.User, error)
}

func NewAuthSQL(db *sql.DB) *AuthSQL {
	return &AuthSQL{
		db: db,
	}
}

func (a *AuthSQL) CreateUser(User model.User) error {
	stmt, err := a.db.Prepare("INSERT INTO User(username, password,email) values(?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(User.Username, User.Password, User.Email); err != nil {
		return err
	}
	return nil
}

func (a *AuthSQL) CheckUser(User model.User) (model.User, error) {
	var fullUser model.User
	query := `SELECT * FROM user WHERE username=$1 and password=$2`
	row := a.db.QueryRow(query, User.Username, User.Password)
	if err := row.Scan(&fullUser.ID, &fullUser.Username, &fullUser.Password, &fullUser.Email); err != nil {
		return fullUser, err
	}
	return fullUser, nil
}

func (a *AuthSQL) CheckUserByToken(token string) (model.User, error) {
	var fullUser model.User
	var id int
	var expiresAt string
	query := `SELECT userId,expiresAt FROM user_sessions WHERE token=$1`
	err := a.db.QueryRow(query, token).Scan(&id, &expiresAt)
	if err != nil {
		return fullUser, err
	}
	query1 := `SELECT userId, username, email FROM user WHERE userId=?`
	row := a.db.QueryRow(query1, id)
	if err := row.Scan(&fullUser.ID, &fullUser.Username, &fullUser.Email); err != nil {
		return fullUser, err
	}
	fullUser.TokenDuration, _ = time.Parse("01-02-2006 15:04:05", expiresAt)
	return fullUser, nil
}

func (a *AuthSQL) SaveToken(User model.User) error {
	stmt, err := a.db.Prepare(`INSERT INTO user_sessions(token, expiresAt,userId) values(?,?,?)`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(User.Token, User.TokenDuration, User.ID); err != nil {
		return err
	}
	return nil
}

func (a *AuthSQL) DeleteToken(token string) error {
	query := `DELETE FROM user_sessions WHERE token=$1`
	_, err := a.db.Exec(query, token)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthSQL) DeleteTokenById(id int) {
	query := `DELETE FROM user_sessions WHERE userId=$1`
	a.db.Exec(query, id)
}

func (a *AuthSQL) GetAvatar(id int) (string, error) {
	var base string
	query1 := `SELECT base FROM user_avatar WHERE userId=?`
	row := a.db.QueryRow(query1, id)
	if err := row.Scan(&base); err != nil {
		return base, err
	}
	return base, nil
}

func (a *AuthSQL) UpdateAvatar(file string, id int) error {
	query := `UPDATE user_avatar SET base = $1 WHERE userId = $4`
	if _, err := a.db.Exec(query, file, id); err != nil {
		return err
	}
	return nil
}

func (a *AuthSQL) UploadAvatar(file string, id int) error {
	stmt, err := a.db.Prepare(`INSERT INTO user_avatar(userId, base) values(?,?)`)
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(id, file); err != nil {
		return err
	}
	return nil
}

func (a *AuthSQL) GetUserInfo(userID int) (model.User, error) {
	var fullUser model.User
	query := `SELECT
		u.userId,
		u.username,
		u.email,
		COALESCE(ua.base, 'null') AS base,
		(SELECT COUNT(id) FROM likesPost WHERE userId = u.userId and like1 =1) as LikedPosts,
		(SELECT COUNT(postId) FROM posts WHERE author = u.username) as Myposts,
		COALESCE((SELECT SUM(like) FROM posts WHERE author = u.username), 0) as LikesOnMyPost
	FROM 
		user as u
	LEFT JOIN 
		user_avatar as ua
	ON 
		ua.userId = u.userId
	WHERE 
		u.userId = $1`
	row := a.db.QueryRow(query, userID)
	if err := row.Scan(&fullUser.ID, &fullUser.Username, &fullUser.Email, &fullUser.Avatar, &fullUser.LikedPosts, &fullUser.Myposts, &fullUser.LikesOnMyPost); err != nil {
		return fullUser, err
	}
	return fullUser, nil
}

func (a *AuthSQL) GetUserByEmail(email string) (model.User, error) {
	query := "SELECT * FROM user WHERE email = $1"
	row := a.db.QueryRow(query, email)

	var fullUser model.User
	err := row.Scan(&fullUser.ID, &fullUser.Username, &fullUser.Password, &fullUser.Email)
	if err != nil {
		return model.User{}, err
	}

	return fullUser, nil
}

func (a *AuthSQL) GetUserByUsername(username string) (model.User, error) {
	query := "SELECT * FROM user WHERE username = $1"
	row := a.db.QueryRow(query, username)

	var fullUser model.User
	err := row.Scan(&fullUser.ID, &fullUser.Username, &fullUser.Password, &fullUser.Email)
	if err != nil {
		return model.User{}, err
	}

	return fullUser, nil
}
