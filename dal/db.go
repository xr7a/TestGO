package dal

import (
	"awesomeProject/config"
	"context"
	"errors"
	"fmt"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

func NewDb(lc fx.Lifecycle, config config.Config) (*sqlx.DB, error) {
	ctx := context.Background()
	log := zap.NewExample()
	db, err := NewPgx(ctx, config)
	if err != nil {
		log.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err = goose.Up(db.DB, "dal/migrations"); err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}

func NewPgx(ctx context.Context, config config.Config) (*sqlx.DB, error) {
	pool, err := newPgxPool(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to pgx pool: %w", err)
	}

	db, err := wrapPgxPool(ctx, pool)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap pgx pool: %w", err)
	}

	return db, nil
}

func newPgxPool(ctx context.Context, config config.Config) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(config.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}

	if poolConfig == nil {
		return nil, errors.New("invalid connection string")
	}

	if config.Traces {
		poolConfig.ConnConfig.Tracer = otelpgx.NewTracer()
	}

	return pgxpool.NewWithConfig(ctx, poolConfig)
}

func wrapPgxPool(ctx context.Context, pool *pgxpool.Pool) (*sqlx.DB, error) {
	db := sqlx.NewDb(stdlib.OpenDBFromPool(pool), "pgx")

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		return nil, fmt.Errorf("failed to ping pgx pool: %w", err)
	}

	return db, nil
}
