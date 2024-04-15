package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	*bun.DB
}

func InitDB(url string) DB {
	if url == "" {
		log.Fatalf("Database URL is not provided")
	}

	db := connectDB(url)
	handler := DB{db}

	if !handler.checkDataExists(&models.Record{}, "journal.yml") {
		handler.ensureFixturesLoaded("journal.yml")
	} else {
		log.Println("Existing data detected, skipping fixtures.")
	}

	return handler
}

func connectDB(url string) *bun.DB {
	const maxRetries = 12
	const retryDelay = 5 * time.Second

	var db *bun.DB
	for i := 0; i < maxRetries; i++ {
		connector := pgdriver.NewConnector(pgdriver.WithDSN(url))
		sqldb := sql.OpenDB(connector)
		tempDB := bun.NewDB(sqldb, pgdialect.New())

		if err := tempDB.Ping(); err != nil {
			log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
			if i < maxRetries-1 {
				log.Printf("Retrying in %v...", retryDelay)
				time.Sleep(retryDelay)
				continue
			} else {
				log.Fatalf("Failed to connect to database after %d attempts, stopping application", maxRetries)
			}
		}

		tempDB.RegisterModel((*models.Record)(nil))
		tempDB.RegisterModel((*models.Journal)(nil))

		db = tempDB
		log.Println("Connected to database successfully.")
		break
	}

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
