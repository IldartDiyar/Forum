package repository

import (
	"database/sql"

	"forum_backend/model"
)

type IEmotionSQL interface {
	CreateEmotionPost(model.Like) error
	GetEmotionPost(model.Like) (model.Like, error)
	EmotionPostExists(postID, userID int) (bool, error)
	UpdateEmotionPost(model.Like) error

	CreateEmotionComment(model.Like) error
	GetEmotionComment(model.Like) (model.Like, error)
	EmotionCommentExists(commentID, userID int) (bool, error)
	UpdateEmotionComment(model.Like) error
}

type EmotionSQL struct {
	db *sql.DB
}

func NewEmotionSQL(db *sql.DB) IEmotionSQL {
	return &EmotionSQL{
		db: db,
	}
}

func (e *EmotionSQL) CreateEmotionPost(PostEmo model.Like) error {
	stmt, err := e.db.Prepare("INSERT INTO likesPost(userId, postId,like1) values(?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(PostEmo.UserID, PostEmo.PostID, PostEmo.Islike); err != nil {
		return err
	}

	asd, err := e.GetEmotionPost(PostEmo)
	if err != nil {
		return err
	}

	query := `UPDATE posts SET like = $1 , dislike  = $2 WHERE postId = $3`
	if _, err := e.db.Exec(query, asd.CountLike, asd.Countdislike, PostEmo.PostID); err != nil {
		return err
	}
	return nil
}

func (e *EmotionSQL) CreateEmotionComment(CommentEmo model.Like) error {
	stmt, err := e.db.Prepare("INSERT INTO likesComment(userId, commentsId,like1) values(?,?,?)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(CommentEmo.UserID, CommentEmo.CommentID, CommentEmo.Islike); err != nil {
		return err
	}

	asd, err := e.GetEmotionComment(CommentEmo)
	if err != nil {
		return err
	}
	query := `UPDATE comments SET like = $1 , dislike = $2 WHERE commentsId = $3`
	if _, err := e.db.Exec(query, asd.CountLike, asd.Countdislike, CommentEmo.CommentID); err != nil {
		return err
	}
	return nil
}

func (e *EmotionSQL) GetEmotionPost(PostEmo model.Like) (model.Like, error) {
	query := `SELECT 
    (SELECT COUNT(*) FROM likesPost WHERE postId = $1 AND like1 = 1) AS likes, 
    (SELECT COUNT(*) FROM likesPost WHERE postId = $1 AND like1 = 0) AS dislikes;`
	row := e.db.QueryRow(query, PostEmo.PostID)
	if err := row.Scan(&PostEmo.CountLike, &PostEmo.Countdislike); err != nil {
		return PostEmo, err
	}
	return PostEmo, nil
}

func (e *EmotionSQL) GetEmotionComment(CommentEmo model.Like) (model.Like, error) {
	query := `SELECT 
    (SELECT COUNT(*) FROM likesComment WHERE commentsId = $1 AND like1 = 1) AS likes, 
    (SELECT COUNT(*) FROM likesComment WHERE commentsId = $1 AND like1 = 0) AS dislikes;`
	row := e.db.QueryRow(query, CommentEmo.CommentID)
	if err := row.Scan(&CommentEmo.CountLike, &CommentEmo.Countdislike); err != nil {
		return CommentEmo, err
	}
	return CommentEmo, nil
}

func (e *EmotionSQL) EmotionCommentExists(commentID, userID int) (bool, error) {
	query := "SELECT COUNT(*) FROM likesComment WHERE commentsId=$1 AND userId=$2"
	var count int
	err := e.db.QueryRow(query, commentID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (e *EmotionSQL) EmotionPostExists(postID, userID int) (bool, error) {
	query := "SELECT COUNT(*) FROM likesPost WHERE postId=$1 AND userId=$2"
	var count int
	err := e.db.QueryRow(query, postID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (e *EmotionSQL) UpdateEmotionComment(CommentEmo model.Like) error {
	query := `UPDATE likesComment SET like1 = $1 WHERE userId = $2 AND commentsId = $3`
	_, err := e.db.Exec(query, CommentEmo.Islike, CommentEmo.UserID, CommentEmo.CommentID)
	if err != nil {
		return err
	}

	asd, err := e.GetEmotionComment(CommentEmo)
	if err != nil {
		return err
	}
	query = `UPDATE comments SET like = $1 , dislike = $2 WHERE commentsId = $3`
	if _, err := e.db.Exec(query, asd.CountLike, asd.Countdislike, CommentEmo.CommentID); err != nil {
		return err
	}
	return nil
}

func (e *EmotionSQL) UpdateEmotionPost(PostEmo model.Like) error {
	query := `UPDATE likesPost SET like1 = $1 WHERE userId = $2 AND postId = $3`
	_, err := e.db.Exec(query, PostEmo.Islike, PostEmo.UserID, PostEmo.PostID)
	if err != nil {
		return err
	}
	asd, err := e.GetEmotionPost(PostEmo)
	if err != nil {
		return err
	}
	query = `UPDATE posts SET like = $1, dislike = $2 WHERE postId = $3`

	if _, err := e.db.Exec(query, asd.CountLike, asd.Countdislike, PostEmo.PostID); err != nil {
		return err
	}
	return nil
}
