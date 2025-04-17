package interceptor

import (
	"1337b0rd/internal/types/controller"
	"context"
	"crypto/rand"
	"fmt"
)

type findUserReqInters struct {
	sessionID string
}
type respAvatar struct {
	name      string
	id        int
	imageUrl  string
	status    bool
	sessionID string
}

func (i *Interceptor) InterceptorGov(ctx context.Context) (controller.RespAvatar, error) {
	resp, err := i.redis.GetAvatarInRedis(ctx)
	if err != nil {
		return nil, err
	}
	newGen, err := i.generateSessionID()
	if err != nil {
		return nil, err
	}
	newRespAvatar := respAvatar{
		name:      resp.GetName(),
		id:        resp.GetID(),
		imageUrl:  resp.GetImageURL(),
		sessionID: newGen,
	}
	return &newRespAvatar, nil
}

func (i *Interceptor) generateSessionID() (string, error) {
	newID := make([]byte, 16)
	_, err := rand.Read(newID)
	if err != nil {
		return "", err
	}
	newID[6] = (newID[6] & 0x0f) | 0x40
	newID[8] = (newID[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		newID[0:4],
		newID[4:6],
		newID[6:8],
		newID[8:10],
		newID[10:16],
	), nil
}

func (a *respAvatar) GetID() int {
	return a.id
}
func (a *respAvatar) GetImageURL() string {
	return a.imageUrl
}
func (a *respAvatar) GetName() string {
	return a.name
}

func (a *respAvatar) GetSessionID() string {
	return a.sessionID
}
