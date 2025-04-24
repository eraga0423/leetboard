package my_redis

import (
	redistypes "1337b0rd/internal/types/redis"
	"context"
	"fmt"
	"strconv"
	"strings"
)

type avatar struct {
	id        int
	name      string
	image     string
	status    bool
	sessionID string
}
type avatars struct {
	allAvatars []avatar
}

func (m MyRedis) SetAvatarsInRedis(avatars redistypes.ReqAvatars, ctx context.Context) error {
	respAvatars := make([]avatar, 0)
	for _, v := range avatars.GetAvatars() {
		respAvatars = append(respAvatars, avatar{
			id:     v.GetID(),
			name:   v.GetName(),
			image:  v.GetImageURL(),
			status: v.GetStatus(),
		})
	}

	for _, v := range respAvatars {
		key := fmt.Sprintf("avatar:%d", v.id)
		err := m.newClient.HSet(ctx, key, map[string]interface{}{
			"name":   v.name,
			"image":  v.image,
			"status": v.status,
		}).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m MyRedis) RefreshAvatarInRedis(ctx context.Context) (redistypes.RespAvatars, error) {
	result := []avatar{}

	var cursor uint64
	for {
		keys, nextCursor, err := m.newClient.Scan(ctx, cursor, "avatar:*", 100).Result()
		if err != nil {
			return nil, err

		}
		for _, v := range keys {
			results, err := m.newClient.HGetAll(ctx, v).Result()
			if err != nil {
				return nil, err

			}
			newID, err := strconv.Atoi(strings.TrimPrefix(v, "avatar:"))
			if err != nil {
				return nil, err
			}
			boolStatus, err := strconv.ParseBool(results["status"])
			if err != nil {
				return nil, err
			}
			result = append(result, avatar{
				id:     newID,
				name:   results["name"],
				image:  results["image"],
				status: boolStatus,
			})

		}
		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	return &avatars{allAvatars: result}, nil
}

func (a *avatar) GetID() int {
	return a.id
}
func (a *avatar) GetName() string {
	return a.name
}
func (a *avatar) GetImageURL() string {
	return a.image
}
func (a *avatar) GetStatus() bool {
	return a.status
}
func (a *avatars) GetAvatars() []redistypes.Avatar {
	resAvatars := make([]redistypes.Avatar, len(a.allAvatars))
	for i, allAvatar := range a.allAvatars {
		resAvatars[i] = &allAvatar
	}
	return resAvatars
}
