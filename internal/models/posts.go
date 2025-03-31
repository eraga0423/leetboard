package models

import "time"

type Post struct {
	PostID      int
	PostTitle   string
	PostContent string
	PostImage   string
	PostTime    time.Time
	User
	Comments
}
