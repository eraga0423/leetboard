package database

type NewPostReq interface {
	GetTitle() string
	GetPostContent() string
	GetImage() string
	GetAuthor() (idUser int)
}

type NewPostResp interface {
	CreationPostResp() (idPost int)
}
