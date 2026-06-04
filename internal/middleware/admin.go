package middleware

import (
    "net/http"

    "github.com/danielgtaylor/huma/v2"
)

func RequireAdmin(api huma.API) func(huma.Context, func(huma.Context)) {

    return func(ctx huma.Context, next func(huma.Context)) {

        role, ok := ctx.Context().Value(RoleKey).(string)

        if !ok {
            huma.WriteErr(
                api,
                ctx,
                http.StatusUnauthorized,
                "Role not found",
            )
            return
        }

        if role != "admin" {
            huma.WriteErr(
                api,
                ctx,
                http.StatusForbidden,
                "Admin access required",
            )
            return
        }

        next(ctx)
    }
}