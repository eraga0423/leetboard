package my_redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
)

func (m MyRedis) GetAvatarInRedis(ctx context.Context) (avatar, error) {
	result := avatar{}
	client := m.newClient()
	var cursor uint64

	for {
		keys, nextCursor, err := client.Scan(ctx, cursor, "avatar:*", 100).Result()
		if err != nil {
			return avatar{}, err
		}

		for _, v := range keys {
			status, err := client.HGet(ctx, v, "status").Result()
			if err != nil {
				return avatar{}, err
			}

			if status == "1" {
				data, err := client.HGetAll(ctx, v).Result()
				if err != nil {
					return avatar{}, err
				}

				err = client.HSet(ctx, v, "status", "0").Err()
				if err != nil {
					return avatar{}, err
				}

				newId, err := strconv.Atoi(strings.TrimPrefix(v, "avatar:"))
				if err != nil {
					return avatar{}, err
				}

				result = avatar{
					id:    newId,
					name:  data["name"],
					image: data["image"],
				}

				return result, nil
			}
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	if err := m.resetAvatars(ctx, client); err != nil {
		return avatar{}, err
	}

	return m.GetAvatarInRedis(ctx)
}

func (m MyRedis) resetAvatars(ctx context.Context, client *redis.Client) error {
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
