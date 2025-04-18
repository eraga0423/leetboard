package interceptor

import (
	"context"
	"log"

	"1337b0rd/internal/types/database"
	redistypes "1337b0rd/internal/types/redis"
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
	log.Print("fetch and cache avatar")
	databaseList, err := i.db.ListCharacters()

	if err != nil {
		log.Print("error in fetch and cache avatar")
		return err

	}
	newList := reqAvatars{}
	if databaseList == nil {
		log.Print("characters are emptyy")
		list, err := i.parseAvatar.ParseDataJson()
		if err != nil {
			return err
		}

		for _, v := range list.RespParseDataJson() {
			newList.avatars = append(newList.avatars, avatar{
				name:     v.GetName(),
				id:       v.GetId(),
				imageURL: v.GetImage(),
				status:   v.GetStatus(),
			})
		}
		err = i.db.InserCartoonCharacters(&newList)
		if err != nil {
			return err
		}
		err = i.redis.SetAvatarsInRedis(&newList, ctx)
		if err != nil {
			return err
		}
		return nil
	}

	list := databaseList.GetCharacters()
	for _, v := range list {
		newList.avatars = append(newList.avatars, avatar{
			name:     v.GetCharacterName(),
			id:       v.GetCharacterId(),
			imageURL: v.GetCharacterImage(),
			status:   v.GetCharacterStatus(),
		})
	}
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

func (r *reqAvatars) Insert() []database.InsertCharacter {
	newAvatars := make([]database.InsertCharacter, len(r.avatars))
	for i, v := range r.avatars {
		newAvatars[i] = &v
	}
	return newAvatars
}

func (r *avatar) InsCharacterId() int {
	return r.id
}

func (r *avatar) InsCharacterName() string {
	return r.name
}

func (r *avatar) InsCharacterImage() string {
	return r.imageURL
}

func (r *avatar) InsCharacterStatus() bool {
	return r.status
}
