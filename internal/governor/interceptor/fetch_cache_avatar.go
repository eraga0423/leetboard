package interceptor

import (
	redistypes "1337b0rd/internal/types/redis"
	"context"
)

type reqAvatars struct {
	avatars []avatar
}

type avatar struct {
	id       int
	name     string
	imageURL string
	status   bool
}

func (i *Interceptor) FetchAndCacheAvatar(ctx context.Context) error {

	//
	//
	//
	//
	//
	///
	//
	//
	//
	/// check database
	//i.redis.GetAvatarInRedis(ctx)

	list, err := i.parseAvatar.ParseDataJson()
	if err != nil {
		return err
	}
	newList := reqAvatars{}
	for _, v := range list.RespParseDataJson() {
		newList.avatars = append(newList.avatars, avatar{
			name:     v.GetName(),
			id:       v.GetId(),
			imageURL: v.GetImage(),
			status:   v.GetStatus(),
		})
	}
	//
	//
	err = i.redis.SetAvatarsInRedis(&newList, ctx)
	if err != nil {
		return err
	}
	return nil
}
func (a *reqAvatars) GetAvatars() []redistypes.Avatar {
	newList := make([]redistypes.Avatar, len(a.avatars))
	for i, a2 := range a.avatars {
		newList[i] = &a2
	}

	return newList
}

func (a *avatar) GetName() string {
	return a.name
}
func (a *avatar) GetImageURL() string {
	return a.imageURL
}
func (a *avatar) GetStatus() bool {
	return a.status
}

func (a *avatar) GetID() int {
	return a.id
}
