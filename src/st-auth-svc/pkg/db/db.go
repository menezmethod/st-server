package db

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"st-auth-svc/pkg/models" // Ensure this import path is correct
)

type Handler struct {
	*bun.DB
}

func Init(url string) Handler {
	if url == "" {
		log.Println("Database URL is not provided")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url)))
	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())

	db.RegisterModel((*models.User)(nil))

	if err := fixture.Load(context.Background(), os.DirFS("./pkg/db"), "user.yml"); err != nil {
		log.Fatalf("Failed to load database fixtures: %v", err)
	}

	log.Println("Database initialized and fixtures loaded successfully")

	return Handler{db}
}
