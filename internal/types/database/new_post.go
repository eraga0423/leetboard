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
	TxRollback(bool) error
	CreationPostResp() (idPost int)
}
