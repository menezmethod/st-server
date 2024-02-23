package db

import (
	"context"
	"database/sql"
	"fmt"
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

func Init(url string) (*DB, error) {
	if url == "" {
		return nil, fmt.Errorf("database URL is not provided")
	}

	db, err := connectAndConfigure(url)
	if err != nil {
		return nil, err
	}

	if err := loadFixturesIfNeeded(db); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func connectAndConfigure(url string) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url)))
	return bun.NewDB(sqldb, pgdialect.New()), nil
}

func loadFixturesIfNeeded(db *bun.DB) error {
	need, err := needToLoadFixtures(db)
	if err != nil {
		return fmt.Errorf("checking fixtures need: %w", err)
	}
	if need {
		return loadFixtures(db)
	}
	fmt.Println("Fixtures already loaded, skipping")
	return nil
}

func needToLoadFixtures(db *bun.DB) (bool, error) {
	count, err := db.NewSelect().Model((*models.Journal)(nil)).Count(context.Background())
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func loadFixtures(db *bun.DB) error {
	db.RegisterModel(&models.Trade{}, &models.Journal{})
	fixture := dbfixture.New(db)
	if err := fixture.Load(context.Background(), os.DirFS("./pkg/db"), "journal.yml"); err != nil {
		return fmt.Errorf("loading fixtures: %w", err)
	}
	fmt.Println("Fixtures loaded successfully")
	return nil
}
