package repository

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"forum_backend/model"
)

type PostSQL struct {
	db *sql.DB
}

type IPostSQL interface {
	CreatePost(model.Post) error
	GetAllPost(int, int) ([]model.Post, error)
	GetPost(int, int) (model.Post, error)
	DeletPost(int) error
	GetMetaDataPost(int) (model.MetaDataPost, error)
	GetMetaDataMyPost(int, int) (model.MetaDataPost, error)
	GetMetaDataMyLikedPost(int, int) (model.MetaDataPost, error)
	GetCategory() ([]model.Category, error)
	GetCategoryID(string) (int, error)
	GetMyPost(int, int, int) ([]model.Post, error)
	GetMyLikedPost(int, int, int) ([]model.Post, error)
	GetUsernameByID(int) (string, error)
}

func NewPostSQL(db *sql.DB) IPostSQL {
	return &PostSQL{
		db: db,
	}
}

const limit = 10

func (p *PostSQL) CreatePost(post model.Post) error {
	query := `INSERT INTO posts(author, title, content, creationDate, category_id, ImageName, ImageBase) 
				VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := p.db.Exec(query, post.Author, post.Title, post.Content, time.Now().Format("01-02-2006 15:04:05"), post.CategoryId, post.ImageName, post.ImageData)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return errors.New("post with that title already exists")
		}
		log.Println(err.Error())
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows were affected")
	}
	return nil
}

func (p *PostSQL) GetUsernameByID(id int) (string, error) {
	var username string
	query := `SELECT username FROM user WHERE userId = $1`
	err := p.db.QueryRow(query, id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (p *PostSQL) GetMyLikedPost(id int, offset int, category int) ([]model.Post, error) {
	var allPosts []model.Post
	var query string
	var args []interface{}

	if category == 0 {
		query = `SELECT
        p.postId,
        p.author,
        p.title,
        p.content,
        p.like,
        p.dislike,
        p.creationDate,
        p.category_id,
        c.category_name,
        COALESCE(ua.base, 'null') AS base,
        COALESCE(p.ImageName, 'null') AS ImageName,
        COALESCE(p.ImageBase, 'null') AS ImageBase
    FROM
        posts p
    JOIN
        posts_category c
    ON
        p.category_id = c.category_id
    JOIN 
        user u
    ON
        p.author = u.username
    LEFT JOIN
        user_avatar ua
    ON
        u.userId = ua.userId
    JOIN
        likesPost lp
    ON
        p.postId = lp.postId
    WHERE
        lp.userId = $1 AND lp.like1 = 1
    LIMIT $2 OFFSET $3`
		args = []interface{}{id, limit, offset}
	} else {
		query = `SELECT
        p.postId,
        p.author,
        p.title,
        p.content,
        p.like,
        p.dislike,
        p.creationDate,
        p.category_id,
        c.category_name,
        COALESCE(ua.base, 'null') AS base,
        COALESCE(p.ImageName, 'null') AS ImageName,
        COALESCE(p.ImageBase, 'null') AS ImageBase
    FROM
        posts p
    JOIN
        posts_category c
    ON
        p.category_id = c.category_id
    JOIN 
        user u
    ON
        p.author = u.username
    LEFT JOIN
        user_avatar ua
    ON
        u.userId = ua.userId
    JOIN
        likesPost lp
    ON
        p.postId = lp.postId
    WHERE
        lp.userId = $1 AND lp.like1 = 1
	AND
		p.category_id = $2
    LIMIT $3 OFFSET $4`
		args = []interface{}{id, category, limit, offset}
	}

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Author, &post.Title, &post.Content, &post.CountLike, &post.CountDislike, &post.CreationDate, &post.CategoryId, &post.Category, &post.AuthorAvatar, &post.ImageName, &post.ImageData); err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}

	return allPosts, nil
}

func (p *PostSQL) GetMyPost(id int, offset int, category int) ([]model.Post, error) {
	var allPosts []model.Post
	var query string
	var args []interface{}

	if category == 0 {
		query = `SELECT
		p.postId,
		p.author,
		p.title,
		p.content,
		p.like,
		p.dislike,
		p.creationDate,
		p.category_id,
		c.category_name,
		COALESCE(ua.base, 'null') AS base,
		COALESCE(p.ImageName, 'null') AS ImageName,
		COALESCE(p.ImageBase, 'null') AS ImageBase
	  FROM
		posts p
	  JOIN
		posts_category c
	  ON
		p.category_id = c.category_id
	  JOIN 
		user u
	  ON
		p.author = u.username
	  LEFT JOIN
		user_avatar ua
	  ON
		u.userId = ua.userId
	  WHERE
		u.userId = $1
		LIMIT $2 OFFSET $3;`
		args = []interface{}{id, limit, offset}
	} else {
		query = `SELECT
		p.postId,
		p.author,
		p.title,
		p.content,
		p.like,
		p.dislike,
		p.creationDate,
		p.category_id,
		c.category_name,
		COALESCE(ua.base, 'null') AS base,
		COALESCE(p.ImageName, 'null') AS ImageName,
		COALESCE(p.ImageBase, 'null') AS ImageBase
	  FROM
		posts p
	  JOIN
		posts_category c
	  ON
		p.category_id = c.category_id
	  JOIN 
		user u
	  ON
		p.author = u.username
	  LEFT JOIN
		user_avatar ua
	  ON
		u.userId = ua.userId
	  WHERE
		u.userId = $1
	  AND
	    p.category_id = $2
		LIMIT $3 OFFSET $4`
		args = []interface{}{id, category, limit, offset}
	}

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Author, &post.Title, &post.Content, &post.CountLike, &post.CountDislike, &post.CreationDate, &post.CategoryId, &post.Category, &post.AuthorAvatar, &post.ImageName, &post.ImageData); err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}

	return allPosts, nil
}

func (p *PostSQL) GetAllPost(offset int, category int) ([]model.Post, error) {
	var allPosts []model.Post
	var query string
	var args []interface{}

	if category == 0 {
		query = `SELECT
		p.postId,
		p.author,
		p.title,
		p.content,
		p.like,
		p.dislike,
		p.creationDate,
		p.category_id,
		c.category_name,
		COALESCE(ua.base, 'null') AS base,
		COALESCE(p.ImageName, 'null') AS ImageName,
		COALESCE(p.ImageBase, 'null') AS ImageBase
		FROM
			posts p
		JOIN
			posts_category c
		ON
			p.category_id = c.category_id
		JOIN 
			user u
		ON
			p.author = u.username
		LEFT JOIN
			user_avatar ua
		ON
			u.userId = ua.userId
		LIMIT $1 OFFSET $2;`
		args = []interface{}{limit, offset}
	} else {
		query = `SELECT
		p.postId,
		p.author,
		p.title,
		p.content,
		p.like,
		p.dislike,
		p.creationDate,
		p.category_id,
		c.category_name,
		COALESCE(ua.base, 'null') AS base,
		COALESCE(p.ImageName, 'null') AS ImageName,
		COALESCE(p.ImageBase, 'null') AS ImageBase
		FROM
			posts p
		JOIN
			posts_category c
		ON
			p.category_id = c.category_id
		JOIN 
			user u
		ON
			p.author = u.username
		LEFT JOIN
			user_avatar ua
		ON
			u.userId = ua.userId
		WHERE 
			p.category_id = $1
			LIMIT $2 OFFSET $3;`
		args = []interface{}{category, limit, offset}
	}

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.Author, &post.Title, &post.Content, &post.CountLike, &post.CountDislike, &post.CreationDate, &post.CategoryId, &post.Category, &post.AuthorAvatar, &post.ImageName, &post.ImageData); err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}

	return allPosts, nil
}

func (p *PostSQL) GetMetaDataPost(category int) (model.MetaDataPost, error) {
	var metadatapost model.MetaDataPost
	if category == 0 {
		if err := p.db.QueryRow("SELECT COUNT(*), ROUND(COUNT(*)/10+0.5) FROM posts").Scan(&metadatapost.Posts, &metadatapost.Pages); err != nil {
			return metadatapost, err
		}
		return metadatapost, nil
	}
	if err := p.db.QueryRow("SELECT COUNT(*), ROUND(COUNT(*)/10+0.5) FROM posts WHERE category_id = $1", category).Scan(&metadatapost.Posts, &metadatapost.Pages); err != nil {
		return metadatapost, err
	}
	return metadatapost, nil
}

func (p *PostSQL) GetMetaDataMyPost(id int, category int) (model.MetaDataPost, error) {
	var metadatapost model.MetaDataPost
	var query string
	var args []interface{}
	if category == 0 {
		query = `SELECT 
		COUNT(*), 
		ROUND(COUNT(*)/10+0.5) 
	FROM 
		posts p 
	WHERE 
		p.author = (SELECT username from user where userId = $1 )`
		args = []interface{}{id}
	} else {
		query = `SELECT 
		COUNT(*), 
		ROUND(COUNT(*)/10+0.5) 
	FROM 
		posts p 
	WHERE 
		p.author = (SELECT username from user where userId = $1 )
		AND
	    p.category_id = $2 
		`
		args = []interface{}{id, category}
	}
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return model.MetaDataPost{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(args...)
	if err := row.Scan(&metadatapost.Posts, &metadatapost.Pages); err != nil {
		return metadatapost, err
	}
	return metadatapost, nil
}

func (p *PostSQL) GetMetaDataMyLikedPost(id int, category int) (model.MetaDataPost, error) {
	var metadatapost model.MetaDataPost
	var query string
	var args []interface{}
	if category == 0 {
		query = `SELECT 
		COUNT(p.postId), 
		ROUND(COUNT(*)/10+0.5) 
	  FROM 
		posts p 
	  JOIN
			likesPost lp
		ON
			p.postId = lp.postId
		WHERE
			lp.userId = $1 AND lp.like1 = 1`
		args = []interface{}{id}
	} else {
		query = `SELECT 
		COUNT(p.postId), 
		ROUND(COUNT(*)/10+0.5) 
	  FROM 
		posts p 
	  JOIN
			likesPost lp
		ON
			p.postId = lp.postId
		WHERE
			lp.userId = $1 AND lp.like1 = 1
		AND
			p.category_id = $2  
		`
		args = []interface{}{id, category}
	}
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return model.MetaDataPost{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(args...)
	if err := row.Scan(&metadatapost.Posts, &metadatapost.Pages); err != nil {
		return metadatapost, err
	}
	return metadatapost, nil
}

func (p *PostSQL) GetPost(postID int, userId int) (model.Post, error) {
	var post model.Post
	query := `SELECT 
		postId,
		author,
		title,
		content,
		like,
		dislike,
		creationDate,
		category_id,
		COALESCE(p.ImageName, 'null') AS ImageName,
		COALESCE(p.ImageBase, 'null') AS ImageBase 
	FROM 
		posts p
	WHERE 
		postid=$1`
	row := p.db.QueryRow(query, postID)
	if err := row.Scan(&post.ID, &post.Author, &post.Title, &post.Content, &post.CountLike, &post.CountDislike, &post.CreationDate, &post.CategoryId, &post.ImageName, &post.ImageData); err != nil {
		return post, err
	}
	if userId != 0 {
		query1 := `SELECT 
		(SELECT EXISTS (SELECT 1 FROM likesPost WHERE userId = $1 AND postId = $2 AND like1 = 1)) AS like, 
		(SELECT EXISTS (SELECT 1 FROM likesPost WHERE userId = $1 AND postId = $2 AND like1 = 0)) AS like;`
		p.db.QueryRow(query1, userId, postID).Scan(&post.Likeisset, &post.Dislikeisset)
	}
	return post, nil
}

func (p *PostSQL) DeletPost(id int) error {
	stmt, err := p.db.Prepare("DELETE FROM posts WHERE postId = $1")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(stmt, id); err != nil {
		return err
	}
	return nil
}

func (p *PostSQL) GetCategoryID(category string) (int, error) {
	var id int
	query := `SELECT category_id FROM posts_category where category_name=$1`
	row := p.db.QueryRow(query, category)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PostSQL) GetCategory() ([]model.Category, error) {
	var allCategories []model.Category
	query := `SELECT * FROM posts_category`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.CategoryID, &category.CategoryName); err != nil {
			return nil, err
		}
		allCategories = append(allCategories, category)
	}
	return allCategories, nil
}
