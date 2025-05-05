package interceptor

import (
	"context"

	"1337b0rd/internal/types/controller"
)

type respAvatars struct {
	allAvatars []oneAvatar
}

type oneAvatar struct {
	name     string
	id       int
	imageURL string
	status   bool
}

func (i *Interceptor) GetAvatarsInRedis(ctx context.Context) (controller.RespAvatars, error) {
	respRedis, err := i.redis.RefreshAvatars(ctx)
	if err != nil {
		return nil, err
	}
	avatars := &respAvatars{}
	list := respRedis.GetAvatars()
	for _, v := range list {
		avatars.allAvatars = append(avatars.allAvatars, oneAvatar{
			name:     v.GetName(),
			id:       v.GetID(),
			imageURL: v.GetImageURL(),
			status:   v.GetStatus(),
		})
	}
	return avatars, nil
}

func (r *respAvatars) GetAvatars() []controller.Avatar {
	newResp := make([]controller.Avatar, len(r.allAvatars))
	for i, v := range r.allAvatars {
		newResp[i] = &v
	}
	return newResp
}

func (r *oneAvatar) GetName() string {
	return r.name
}

func (r *oneAvatar) GetID() int {
	return r.id
}

func (r *oneAvatar) GetImageURL() string {
	return r.imageURL
}

func (r *oneAvatar) GetStatus() bool {
	return r.status
}
