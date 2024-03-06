package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"st-gateway/middleware"
	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
	"st-gateway/pkg/journal"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed loading config: %v\n", err)
		return
	}

	r := setupRouter(cfg)
	if err := r.Run(cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}

func setupRouter(cfg config.Config) *gin.Engine {
	r := gin.Default()
	allowedOrigins := []string{"*"}
	r.Use(middleware.CORS(allowedOrigins))

	authSvc := *auth.RegisterAuthRoutes(r, &cfg)
	journal.RegisterJournalRoutes(r, &cfg, &authSvc)

	return r
}
