package database

type FindUserReq interface {
	GetSessionID() string
}

type FindUserResp interface {
	GetUserName() string
	GetUserImageURL() string
}
