package database

type NewPostReq interface {
	GetTitle() string
	GetPostContent() string
	GetImage() []byte
	GetAuthor() (idUser string)
}

type NewPostResp interface {
	CreationPostResp() (idPost int)
}
