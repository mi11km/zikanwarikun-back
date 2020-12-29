package auth

import (
	"context"
	"net/http"

	database "github.com/mi11km/zikanwarikun-back/internal/db"
	"github.com/mi11km/zikanwarikun-back/internal/db/models"
	"github.com/mi11km/zikanwarikun-back/pkg/jwt"
)

type Auth struct {
	User  *models.User
	Token *string
}

var (
	authCtxKey = &contextKey{"auth"}
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
			var user models.User
			result := database.Db.Where("id = ?", id).First(&user)
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

// GetAuthInfoFromCtx finds the user and token from the context. REQUIRES Middleware to have run.
func GetAuthInfoFromCtx(ctx context.Context) *Auth {
	raw, _ := ctx.Value(authCtxKey).(*Auth)
	return raw
}
