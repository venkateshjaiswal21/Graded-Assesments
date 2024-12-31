package models

import "time"

type BlogPost struct {
	PostID       int       `json:"postId"`
	Headline     string    `json:"headline"`
	ArticleText  string    `json:"articleText"`
	Contributor  string    `json:"contributor"`
	CreatedDate  time.Time `json:"createdDate"`
	LastModified time.Time `json:"lastModified"`
}
