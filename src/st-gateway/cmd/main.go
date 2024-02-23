package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"st-gateway/middlewares"
	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
	"st-gateway/pkg/journal"
)

func main() {
	cfg, err := config.LoadConfig() // Assuming LoadConfig returns a pointer to Config
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := setupRouter(cfg)
	if err := r.Run(cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func setupRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORS(cfg.AllowedOrigins))

	authSvc := auth.RegisterAuthRoutes(r, cfg)
	journal.RegisterJournalRoutes(r, cfg, authSvc)

	return r
}
