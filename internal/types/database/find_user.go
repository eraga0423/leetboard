package database

type FindUserReq interface {
	GetSesionID() int
}
type FindUserResp interface {
	GetUserName() string
	GetUserAvatar() string
	GetUserImageURL() string
}
