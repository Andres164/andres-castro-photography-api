package middleware

import "github.com/gin-gonic/gin"

func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header(
			"Strict-Transport-Security",
			"max-age=31536000; includeSubDomains",
		)

		c.Header(
			"X-Content-Type-Options",
			"nosniff",
		)

		// TODO: Implement? Content-Security-Policy (https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Content-Security-Policy)
		c.Header(
			"X-Frame-Options",
			"DENY",
		)

		c.Next()
	}
}