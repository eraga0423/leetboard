package my_redis

import (
	"1337b0rd/internal/types/controller"
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
)

func (m MyRedis) GetAvatarInRedis(ctx context.Context) (controller.RespAvatar, error) {
	result := avatar{}

	var cursor uint64

	for {
		keys, nextCursor, err := m.newClient.Scan(ctx, cursor, "avatar:*", 100).Result()
		if err != nil {
			return nil, err
		}

		for _, v := range keys {
			status, err := m.newClient.HGet(ctx, v, "status").Result()
			if err != nil {
				return nil, err
			}

			if status == "1" {
				data, err := m.newClient.HGetAll(ctx, v).Result()
				if err != nil {
					return nil, err
				}

				err = m.newClient.HSet(ctx, v, "status", "0").Err()
				if err != nil {
					return nil, err
				}

				newId, err := strconv.Atoi(strings.TrimPrefix(v, "avatar:"))
				if err != nil {
					return nil, err
				}

				result = avatar{
					id:    newId,
					name:  data["name"],
					image: data["image"],
				}

				return &result, nil
			}
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	if err := m.resetAvatars(ctx, m.newClient); err != nil {
		return nil, err
	}

	return m.GetAvatarInRedis(ctx)
}

func (m *MyRedis) resetAvatars(ctx context.Context, client *redis.Client) error {
	var cursor uint64
	for {
		keys, NextCursor, err := client.Scan(ctx, cursor, "avatar:*", 100).Result()
		if err != nil {
			return err
		}
		for _, v := range keys {
			err := client.HSet(ctx, v, "status", "1").Err()
			if err != nil {
				return err
			}
		}
		if NextCursor == 0 {
			break
		}
		cursor = NextCursor
	}
	return nil
}

func (m *avatar) GetSessionID() string {
	return m.sessionID
}
