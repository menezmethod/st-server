package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
	"st-gateway/pkg/journal"
)

func CORS(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		var isAllowed bool
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			c.Header("Access-Control-Allow-Origin", origin)
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

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed loading config: %v\n", err)
		return
	}

	r := gin.Default()
	allowedOrigins := []string{"*"}
	r.Use(CORS(allowedOrigins))

	authSvc := *auth.RegisterAuthRoutes(r, &config)
	journal.RegisterJournalRoutes(r, &config, &authSvc)

	if err := r.Run(config.Port); err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}
