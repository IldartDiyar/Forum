package repository

import (
	"database/sql"
	"forum_backend/model"
	"time"
)

type CommentSQL struct {
	db *sql.DB
}

type ICommentSQL interface {
	CreateComment(model.Comment) error
	GetComment(int, int) ([]model.Comment, error)
	GetUsernameById(int) (string, error)
}

func NewCommentSQL(db *sql.DB) ICommentSQL {
	return &CommentSQL{
		db: db,
	}
}

func (c *CommentSQL) GetUsernameById(id int) (string, error) {
	var username string
	query := `SELECT username FROM user WHERE userId = $1`
	row := c.db.QueryRow(query, id)
	if err := row.Scan(&username); err != nil {
		return username, err
	}
	return username, nil
}

func (c *CommentSQL) CreateComment(comment model.Comment) error {
	stmt, err := c.db.Prepare("INSERT INTO comments(postId, author,content,creationDate) values(?,?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(comment.PostId, comment.Author, comment.Body, time.Now().Format("01-02-2006 15:04:05")); err != nil {
		return err
	}
	return nil
}

func (c *CommentSQL) GetComment(postID int, userID int) ([]model.Comment, error) {
	comments := []model.Comment{}

	stmt, err := c.db.Prepare(`SELECT * FROM comments WHERE PostId=$1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment
		if err := rows.Scan(&comment.ID, &comment.PostId, &comment.Author, &comment.Body, &comment.CountLike, &comment.CountDislike, &comment.CreationDate); err != nil {
			return nil, err
		}

		if userID != 0 {
			if err := c.setLikeAndDislike(userID, &comment); err != nil {
				return nil, err
			}
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (c *CommentSQL) setLikeAndDislike(userID int, comment *model.Comment) error {
	query := `SELECT 
		(SELECT EXISTS (SELECT 1 FROM likesComment WHERE userId = $1 AND commentsId = $2 AND like1 = 1)) AS like, 
		(SELECT EXISTS (SELECT 1 FROM likesComment WHERE userId = $1 AND commentsId = $2 AND like1 = 0)) AS dislike;`

	if err := c.db.QueryRow(query, userID, comment.ID).Scan(&comment.Likeisset, &comment.Dislikeisset); err != nil {
		return err
	}

	return nil
}
