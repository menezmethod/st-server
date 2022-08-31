package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
	"st-gateway/pkg/trade"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed loading config", err)
	}
	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	trade.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
