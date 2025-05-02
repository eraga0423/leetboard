package middleware

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"1337b0rd/internal/constants"
)

type respAvatar struct {
	name      string
	imageURL  string
	id        int
	sessionID string
}

func (m *Middleware) Authentificator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("method", "auth", "")
		defer func() {
			if err := recover(); err != nil {
				slog.Error(
					"panic recovered",
					"method", r.Method,
					"url", r.URL.String(),
					"panic", err,
				)
				http.Error(w, "internal server error", http.StatusInternalServerError)

			}
		}()
		cookie, err := r.Cookie("session_id")
		ctx := r.Context()
		if err != nil || cookie.Value == "" {
			newAvatar, err := m.ctrl.InterceptorGov(ctx)
			if err != nil {
				slog.Error("middleware", "error", err)
				return
			}
			newRespAvatar := respAvatar{
				name:      newAvatar.GetName(),
				imageURL:  newAvatar.GetImageURL(),
				id:        newAvatar.GetID(),
				sessionID: newAvatar.GetSessionID(),
			}
			// if err != nil {
			// 	http.Error(w, "error generate id", http.StatusInternalServerError)
			// 	return
			// }
			cookie = &http.Cookie{
				Name:     constants.SessionIDKey,
				Value:    newRespAvatar.sessionID,
				Path:     "/",
				HttpOnly: true,
				Expires:  time.Now().Add(7 * 24 * time.Hour),
			}
			http.SetCookie(w, cookie)
			cookieName := &http.Cookie{
				Name:     constants.Name,
				Value:    newRespAvatar.name,
				Path:     "/",
				HttpOnly: true,
				Expires:  time.Now().Add(7 * 24 * time.Hour),
			}
			http.SetCookie(w, cookieName)
			cookieImageURL := &http.Cookie{
				Name:     constants.ImageURL,
				Value:    newRespAvatar.imageURL,
				Path:     "/",
				HttpOnly: true,
				Expires:  time.Now().Add(7 * 24 * time.Hour),
			}
			http.SetCookie(w, cookieImageURL)
			cookieID := &http.Cookie{
				Name:     constants.AvatarID,
				Value:    strconv.Itoa(newRespAvatar.id),
				Path:     "/",
				HttpOnly: true,
				Expires:  time.Now().Add(7 * 24 * time.Hour),
			}
			http.SetCookie(w, cookieID)
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Incoming HTTP request",
			"protocol", r.Proto, // Пример: "HTTP/1.1"
			"method", r.Method, // Пример: "GET", "POST"
			"uri", r.RequestURI, // Пример: "/post/123?reply=1"
		)
		next.ServeHTTP(w, r)
	})
}
