package models

type UserMetadata struct {
}
type User struct {
	UserID     int
	UserName   string
	UserAvatar string
	UserMetadata
}
