package middleware

import (
	"context"

	"1337b0rd/internal/constants"
)

func withSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, constants.SessionIDKey, sessionID)
}
