package interceptor

import (
	"context"
	"crypto/rand"
	"fmt"
	"strings"

	"1337b0rd/internal/types/controller"
)

type respAvatar struct {
	name      string
	id        int
	imageUrl  string
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
	sessionID := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		newID[0:4],
		newID[4:6],
		newID[6:8],
		newID[8:10],
		newID[10:16],
	)

	sessionID = strings.ToLower(sessionID)
	sessionID = strings.ReplaceAll(sessionID, "-", "")
	sessionID = strings.Trim(sessionID, ".-")

	if len(sessionID) < 3 {
		return "", fmt.Errorf("session ID is too short")
	}
	if len(sessionID) > 63 {
		sessionID = sessionID[:63]
	}

	return sessionID, nil
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
