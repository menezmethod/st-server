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

	if err := prepareDatabase(db); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func connectAndConfigure(url string) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url)))
	bunDB := bun.NewDB(sqldb, pgdialect.New())
	if err := bunDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return bunDB, nil
}

func prepareDatabase(db *bun.DB) error {
	ctx := context.Background()
	exists, err := checkUsersExist(db, ctx)
	if err != nil {
		return fmt.Errorf("failed to check users' existence: %v", err)
	}
	if !exists {
		if err := loadFixtures(db, ctx); err != nil {
			return err
		}
		log.Println("Database fixtures loaded successfully")
	} else {
		log.Println("Database fixtures already exist, skipping loading")
	}
	return nil
}

func checkUsersExist(db *bun.DB, ctx context.Context) (bool, error) {
	count, err := db.NewSelect().Model((*models.User)(nil)).Count(ctx)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func loadFixtures(db *bun.DB, ctx context.Context) error {
	fixture := dbfixture.New(db)
	db.RegisterModel((*models.User)(nil))
	if err := fixture.Load(ctx, os.DirFS("./pkg/db"), "user.yml"); err != nil {
		return fmt.Errorf("failed to load database fixtures: %v", err)
	}
	return nil
}
