package database

type NewPostReq interface {
	GetTitle() string
	GetPostContent() string
	GetPostImageURL() string
	GetAuthorAvatarURL() string
	GetAuthorName() string
	GetAuthorSession() (idSessionUser string)
}

type NewPostResp interface {
	TxRollback() error
	CreationPostResp() (idPost int)
}
