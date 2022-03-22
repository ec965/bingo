package middleware

import (
	"net/http"

	"github.com/ec965/bingo/pkgs/entities"
	"github.com/ec965/bingo/pkgs/token"
)

// Authentication closure style middleware
// example:
// 	authMiddleware := middlware.AuthUser(tokenManager)
// 	r.HandleFunc("/private", authRoute(handlerWithUser))
func AuthUser(tm *token.TokenManager) func(func(http.ResponseWriter, *http.Request, *entities.User)) http.HandlerFunc {
	return func(fn func(http.ResponseWriter, *http.Request, *entities.User)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")
			if tokenStr == "" {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			} else {
				tokenStr = tokenStr[7 : len(tokenStr)-1]
				user, err := tm.ValidateToken(tokenStr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusForbidden)
					return
				}
				fn(w, r, user)
			}
		}
	}
}
