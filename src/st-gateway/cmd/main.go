package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"st-gateway/pkg/journal"

	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed loading config", err)
	}
	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	journal.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
