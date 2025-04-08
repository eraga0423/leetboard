package interceptor

import (
	"context"
	"crypto/rand"
	"fmt"

	"1337b0rd/internal/constants"
)

func (i *Interceptor) InterceptorGov(ctx context.Context, sessionID int) context.Context {
	return context.WithValue(ctx, constants.SessionIDKey, sessionID)
}

func (m *Interceptor) GenerateSessionID() (string, error) {
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
