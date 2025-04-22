package controller

type RespAvatar interface {
	GetName() string
	GetImageURL() string
	GetID() int
	GetSessionID() string

}

type RespAvatars interface {
	GetAvatars() []Avatar
}

type Avatar interface {
	GetName() string
	GetID() int
	GetImageURL() string
	GetStatus() bool
}
