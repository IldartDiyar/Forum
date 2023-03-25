package repository

import (
	"database/sql"
)

type Repository struct {
	IAuthSQL
	IPostSQL
	ICommentSQL
	IEmotionSQL
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		IAuthSQL:    NewAuthSQL(db),
		IPostSQL:    NewPostSQL(db),
		ICommentSQL: NewCommentSQL(db),
		IEmotionSQL: NewEmotionSQL(db),
	}
}
