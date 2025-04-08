package database

type FindUserReq interface {
	GetSessionID() int
}
type FindUserResp interface {
	UserRespFunc() (UserResp, error)
}

type UserResp interface {
	GetUserName() string
	GetUserAvatar() string
	GetUserImageURL() string
}
