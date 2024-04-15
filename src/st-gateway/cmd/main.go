package main

import (
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/helper"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed loading configs: %v\n", err)
		return
	}

	r := setupRouter(cfg)
	if err := r.Run(cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}
}

func setupRouter(cfg configs.Config) *gin.Engine {
	r := gin.Default()
	allowedOrigins := []string{"*"}
	r.Use(auth.CORS(allowedOrigins))

	authSvc := *auth.RegisterAuthRoutes(r, &cfg)
	journal.RegisterJournalRoutes(r, &cfg, &authSvc)
	record.RegisterRecordRoutes(r, &cfg, &authSvc)
	helper.RegisterHelperRoutes(r, &cfg)

	r.GET("/metrics", prometheusHandler())

	return r
}

func prometheusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := promhttp.Handler()
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
