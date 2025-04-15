package interceptor

type RespAvatars struct {
	allAvatars []oneAvatar
}

type oneAvatar struct {
	name     string
	id       int
	imageURL string
}

//func (i *Interceptor) GetAvatarsInRedis(ctx context.Context) (controller.RespAvatars, error) {
//	respRedis, err := i.redis.RefreshAvatars(ctx)
//	if err != nil {
//		return nil, err
//	}
//	avatars := &RespAvatars{}
//	list := respRedis.GetAvatars()
//	for _, v := range list {
//		avatars.allAvatars = append(avatars.allAvatars, oneAvatar{
//			name:     v.GetName(),
//			id:       v.GetID(),
//			imageURL: v.GetImageURL(),
//		})
//	}
//	return , nil
//}
//func ()
