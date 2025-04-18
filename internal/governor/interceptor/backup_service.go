package interceptor

import (
	"1337b0rd/internal/types/database"
	"context"
)

type redisAvatars struct {
	avatars []redisAvatar
}
type redisAvatar struct {
	id     int
	status bool
}

func (i *Interceptor) BackupAvatars(ctx context.Context) error {
	// time.Sleep(5 * time.Minute)
	redisResp, err := i.redis.RefreshAvatars(ctx)
	if err != nil {
		return err
	}
	newRedisAvatars := redisAvatars{}
	listRedis := redisResp.GetAvatars()
	for _, v := range listRedis {
		newRedisAvatars.avatars = append(newRedisAvatars.avatars, redisAvatar{
			id:     v.GetID(),
			status: v.GetStatus(),
		})
	}
	err = i.db.UpdateCharacters(&newRedisAvatars)
	if err != nil {
		return err
	}
	return nil
}

func (r *redisAvatars) SetCharacters() []database.SetCharacter {
	newCharacters := make([]database.SetCharacter, len(r.avatars))
	for i, v := range r.avatars {
		newCharacters[i] = &v
	}
	return newCharacters
}

func (r *redisAvatar) SetCharacterId() int {
	return r.id
}
func (r *redisAvatar) SetCharacterStatus() bool {
	return r.status
}
