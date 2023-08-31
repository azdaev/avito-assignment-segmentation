package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"segmentation-avito/internal/api"
	"segmentation-avito/internal/api/service/segments"
	"segmentation-avito/internal/config"
	"segmentation-avito/internal/pkg/db"
	"segmentation-avito/internal/repository/postgres"
)

func main() {
	cfg := config.MustLoad()

	ctx := context.Background()

	pool, err := db.OpenDB(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	r := api.New(
		segments.New(
			postgres.New(pool),
		),
		logger,
	)

	log.Fatal(r.Run(cfg.ServerPort))
}
