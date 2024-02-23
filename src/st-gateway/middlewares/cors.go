package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORS(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		var isAllowed bool
		if len(allowedOrigins) == 1 && allowedOrigins[0] == "*" {
			isAllowed = true
			c.Header("Access-Control-Allow-Origin", "*")
		} else {
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					isAllowed = true
					c.Header("Access-Control-Allow-Origin", origin)
					break
				}
			}
		}

		if isAllowed {
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
