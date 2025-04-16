package redis_types

import "context"

type TypesRedis interface {
	SetAvatarsRedis(ReqAvatars, context.Context) error
	GetAvatarInRedis(ctx context.Context) (RespAvatar, error)
	RefreshAvatars(ctx context.Context) (RespAvatar, error)
}
type ReqAvatars interface {
	GetAvatars() []Avatar
}
type Avatar interface {
	GetID() int
	GetName() string
	GetImageURL() string
	GetStatus() bool
}

type RespAvatar interface {
	GetID() int
	GetName() string
	GetImageURL() string
	GetStatus() bool
}

type RespAvatars interface {
	GetAvatars() []Avatar
}
