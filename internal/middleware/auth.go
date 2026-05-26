package middleware

import (
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("super-secret-key")

func AuthMiddleware(ctx huma.Context, next func(huma.Context)) {

	authHeader := ctx.Header("Authorization")

	if authHeader == "" {
		huma.WriteErr(ctx, huma.Error401Unauthorized("Missing Authorization header"))
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		huma.WriteErr(ctx, huma.Error401Unauthorized("Invalid Authorization format"))
		return
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		huma.WriteErr(ctx, huma.Error401Unauthorized("Invalid token"))
		return
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		// Store in context
		ctx.Set("user_id", claims["user_id"])
		ctx.Set("role", claims["role"])
	}

	next(ctx)
}