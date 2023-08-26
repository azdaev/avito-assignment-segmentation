package db

import (
	"context"
	"fmt"
	"segmentation-avito/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDB(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("OpenDB config parse: %w", err)
	}

	config.ConnConfig.Host = cfg.PostgersConfig.Host
	config.ConnConfig.Port = cfg.PostgersConfig.Port
	config.ConnConfig.Database = cfg.PostgersConfig.Database
	config.ConnConfig.User = cfg.PostgersConfig.User
	config.ConnConfig.Password = cfg.PostgersConfig.Password

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("OpenDB connect: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("OpenDB ping: %w", err)
	}

	return pool, nil
}
