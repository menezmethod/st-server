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
	"st-auth-svc/pkg/models"
)

type Handler struct {
	*bun.DB
}

func Init(url string) Handler {
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(url))), pgdialect.New())
	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())

	db.RegisterModel(&models.User{})

	if err := fixture.Load(context.Background(), os.DirFS("./pkg/db"), "user.yml"); err != nil {
		log.Fatalln(err)
	}

	return Handler{db}
}
