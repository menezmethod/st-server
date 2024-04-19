package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
)

type DB struct {
	*bun.DB
	logger *zap.Logger
}

func InitDB(url string, logger *zap.Logger) DB {
	if url == "" {
		logger.Fatal("Database URL is not provided")
	}

	db := connectDB(url, logger)
	handler := DB{DB: db, logger: logger}

	if !handler.checkUserDataExists() {
		handler.ensureFixturesLoaded(db)
	} else {
		logger.Info("Existing user data detected, skipping fixtures.")
	}

	return handler
}

func connectDB(url string, logger *zap.Logger) *bun.DB {
	const maxRetries = 12
	const retryDelay = 5 * time.Second

	var db *bun.DB
	for i := 0; i < maxRetries; i++ {
		connector := pgdriver.NewConnector(pgdriver.WithDSN(url))
		sqldb := sql.OpenDB(connector)
		tempDB := bun.NewDB(sqldb, pgdialect.New())

		if err := tempDB.Ping(); err != nil {
			logger.Error("Failed to connect to database", zap.Int("attempt", i+1), zap.Int("max_retries", maxRetries), zap.Error(err))
			if i < maxRetries-1 {
				logger.Info("Retrying database connection", zap.Duration("retry_delay", retryDelay))
				time.Sleep(retryDelay)
				continue
			} else {
				logger.Fatal("Failed to connect to database after max attempts", zap.Int("attempts", maxRetries))
			}
		}

		db = tempDB
		logger.Info("Connected to database successfully.")
		break
	}

	return db
}

func (h DB) ensureFixturesLoaded(db *bun.DB) {
	if err := h.loadFixtures(db); err != nil {
		h.logger.Fatal("Failed to load database fixtures", zap.Error(err))
	}
	h.logger.Info("Database fixtures loaded successfully")
}

func (h DB) checkUserDataExists() bool {
	var count int
	count, err := h.NewSelect().Model((*models.User)(nil)).Limit(1).Count(context.Background())
	if err != nil {
		h.logger.Error("Failed to check for existing user data", zap.Error(err))
		return false
	}
	return count > 0
}

func (_ DB) loadFixtures(db *bun.DB) error {
	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	err := fixture.Load(context.Background(), os.DirFS("./pkg/db"), "user.yml")
	if err != nil {
		return fmt.Errorf("failed to load database fixtures: %w", err)
	}
	return nil
}
