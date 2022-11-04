package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"st-gateway/pkg/journal"

	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
)

func CORS(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed loading config", err)
	}
	r := gin.Default()
	r.Use(CORS)
	//r.Use(cors.Default())
	authSvc := *auth.RegisterAuthRoutes(r, &c)
	journal.RegisterJournalRoutes(r, &c, &authSvc)
	r.Run(c.Port)
}
