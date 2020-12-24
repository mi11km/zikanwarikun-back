package auth

import (
	"context"
	"net/http"

	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"github.com/mi11km/zikanwarikun-back/internal/db/models/users"
	"github.com/mi11km/zikanwarikun-back/pkg/jwt"
)

type Auth struct {
	User  *users.User
	Token *string
}

var (
	authCtxKey  = &contextKey{"auth"}
)

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			id, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			var user users.User
			result := database.Db.Select("id", "email", "password", "school", "name").Where("id = ?", id).First(&user)
			if result.Error != nil {
				next.ServeHTTP(w, r)
				return
			}

			auth := &Auth{&user, &tokenStr}

			// put it in context
			ctx := context.WithValue(r.Context(), authCtxKey, auth)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *Auth {
	raw, _ := ctx.Value(authCtxKey).(*Auth)
	return raw
}


