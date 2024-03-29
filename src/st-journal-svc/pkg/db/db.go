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
	"st-journal-svc/pkg/models"
)

type DB struct {
	*bun.DB
}

func Init(url string) DB {
	if url == "" {
		log.Fatalf("Database URL is not provided")
	}

	db := connectDB(url)
	handler := DB{db}

	if !handler.checkDataExists(&models.Trade{}, "journal.yml") {
		handler.ensureFixturesLoaded("journal.yml")
	} else {
		log.Println("Existing data detected, skipping fixtures.")
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

	db.RegisterModel((*models.Trade)(nil))
	db.RegisterModel((*models.Journal)(nil))

	return db
}

func (h DB) ensureFixturesLoaded(fixtureFile string) {
	if err := h.loadFixtures(fixtureFile); err != nil {
		log.Fatalf("Failed to load database fixtures: %v", err)
	}
	log.Println("Database fixtures loaded successfully")
}

func (h DB) checkDataExists(model interface{}, fixtureFile string) bool {
	count, err := h.NewSelect().Model(model).Limit(1).Count(context.Background())
	if err != nil {
		log.Printf("Failed to check for existing data: %v", err)
		return false
	}
	return count > 0
}

func (h DB) loadFixtures(fixtureFile string) error {
	fixture := dbfixture.New(h.DB, dbfixture.WithRecreateTables())
	err := fixture.Load(context.Background(), os.DirFS("./pkg/db"), fixtureFile)
	if err != nil {
		return fmt.Errorf("failed to load database fixtures for %s: %v", fixtureFile, err)
	}
	return nil
}
