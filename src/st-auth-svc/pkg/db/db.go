package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"st-auth-svc/pkg/models"
)

type Handler struct {
	*bun.DB
}

func InitDB(url string) Handler {
	if url == "" {
		log.Fatalf("Database URL is not provided")
	}

	db := connectDB(url)
	handler := Handler{db}

	if !handler.checkUserDataExists() {
		handler.ensureFixturesLoaded(db)
	} else {
		log.Println("Existing user data detected, skipping fixtures.")
	}

	return handler
}

func connectDB(url string) *bun.DB {
	connector := pgdriver.NewConnector(pgdriver.WithDSN(url))
	sqldb := sql.OpenDB(connector)
	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}

func (h Handler) ensureFixturesLoaded(db *bun.DB) {
	if err := h.loadFixtures(db); err != nil {
		log.Fatalf("Failed to load database fixtures: %v", err)
	}
	log.Println("Database fixtures loaded successfully")
}

func (h Handler) checkUserDataExists() bool {
	var count int
	count, err := h.NewSelect().Model((*models.User)(nil)).Limit(1).Count(context.Background())
	if err != nil {
		log.Printf("Failed to check for existing user data: %v", err)
		return false
	}
	return count > 0
}

func (h Handler) loadFixtures(db *bun.DB) error {
	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	err := fixture.Load(context.Background(), os.DirFS("./pkg/db"), "user.yml")
	if err != nil {
		return fmt.Errorf("failed to load database fixtures: %v", err)
	}
	return nil
}
