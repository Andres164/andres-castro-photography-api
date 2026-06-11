package middleware

import (
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"

	"andres_castro_photography_api/internal/utils"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	RoleKey   contextKey = "role"
)

func AuthMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {

	return func(ctx huma.Context, next func(huma.Context)) {

		authHeader := ctx.Header("Authorization")

		if authHeader == "" {
			huma.WriteErr(api, ctx, 401, "Missing Authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			huma.WriteErr(api, ctx, 401, "Invalid Authorization format")
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return utils.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			huma.WriteErr(api, ctx, 401, "Invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			huma.WriteErr(api, ctx, 401, "Invalid token claims")
			return
		}

		// Add values to request context
		ctx = huma.WithValue(ctx, UserIDKey, claims["user_id"])
		ctx = huma.WithValue(ctx, RoleKey, claims["role"])

		next(ctx)
	}
}