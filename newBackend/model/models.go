package model

import "time"

type User struct {
	ID            int       `json:"Id"`
	Email         string    `json:"Email,omitempty"`
	Username      string    `json:"Username,omitempty"`
	Password      string    `json:"Password,omitempty"`
	Token         string    `json:"Token,omitempty"`
	TokenDuration time.Time `json:"TokenDuration,omitempty"`
	Avatar        string    `json:"Avatar,omitempty"`
	LikedPosts    int       `json:"LikedPosts"`
	Myposts       int       `json:"MyPosts"`
	LikesOnMyPost int       `json:"LikesOnMyPost"`
}

type Post struct {
	ID           int    `json:"Id"`
	Author       string `json:"Author,omitempty"`
	AuthorAvatar string `json:"AuthorAvatar,omitempty"`
	Title        string `json:"Title,omitempty"`
	Content      string `json:"Content,omitempty"`
	CreationDate string `json:"CreationDate,omitempty"`
	CountLike    string `json:"CountLike,omitempty"`
	CountDislike string `json:"CountDislike,omitempty"`
	Dislikeisset bool   `json:"Dislikeisset,omitempty"`
	Likeisset    bool   `json:"Likeisset,omitempty"`
	Category     string `json:"Category,omitempty"`
	CategoryId   int    `json:"CategoryID"`
	ImageName    string `json:"ImageName,omitempty"`
	ImageData    string `json:"ImageData,omitempty"`
}

type Comment struct {
	ID           int    `json:"Id"`
	Author       string `json:"Author,omitempty"`
	Body         string `json:"Body,omitempty"`
	PostId       int    `json:"PostId"`
	CreationDate string `json:"CreationDate,omitempty"`
	Dislikeisset bool   `json:"Dislikeisset,omitempty"`
	Likeisset    bool   `json:"Likeisset,omitempty"`
	CountLike    string `json:"CountLike,omitempty"`
	CountDislike string `json:"CountDislike,omitempty"`
}

type Like struct {
	UserID       int `json:"UserID"`
	PostID       int `json:"PostID"`
	Islike       int `json:"Islike"`
	CommentID    int `json:"CommentID"`
	CountLike    int `json:"CountLike"`
	Countdislike int `json:"Countdislike"`
}

type Category struct {
	CategoryID   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}
type MetaDataPost struct {
	Posts int `json:"Posts"`
	Pages int `json:"Pages"`
}
