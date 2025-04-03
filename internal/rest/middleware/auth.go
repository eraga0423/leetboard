package middleware

import (
	"net/http"
	"time"
)

func (m *Middleware) Authentificator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		ctx := r.Context()
		if err != nil || cookie.Value == "" {
			newSessionID, err := m.ctrl.GenerateSessionID()
			if err != nil {
				http.Error(w, "error generate id", http.StatusInternalServerError)
				return
			}
			cookie = &http.Cookie{
				Name:     "session_id",
				Value:    newSessionID,
				Path:     "/",
				HttpOnly: true,
				Expires:  time.Now().Add(7 * 24 * time.Hour),
			}
			http.SetCookie(w, cookie)
		}
		newContext := m.ctrl.InterceptorGov(ctx, cookie.Value)
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}
