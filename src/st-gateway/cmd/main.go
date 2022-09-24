package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"st-gateway/pkg/journal"

	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed loading config", err)
	}
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(cors.Default())
	authSvc := *auth.RegisterRoutes(r, &c)
	journal.RegisterRoutes(r, &c, &authSvc)
	r.Run(c.Port)
}
