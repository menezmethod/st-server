package db

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"os"
	"st-journal-svc/pkg/models"
)

type DB struct {
	*bun.DB
}

func Init(url string) DB {
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url))), pgdialect.New())
	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())

	db.RegisterModel(&models.Trade{})

	if err := fixture.Load(context.Background(), os.DirFS("./pkg/db"), "trade.yml"); err != nil {
		log.Fatalln(err)
	}

	return DB{db}
}
