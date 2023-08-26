package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	pool *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repo {
	return &Repo{
		pool: db,
	}
}

func (r *Repo) GetUserSegments(ctx context.Context, userID int) ([]string, error) {
	rows, err := r.pool.Query(ctx, "SELECT segment_name FROM users_segments WHERE user_id = $1", userID)
	if err != nil {
		return nil, fmt.Errorf("SQL: get segments: %w", err)
	}

	var segmentsCount int
	r.pool.QueryRow(ctx, "SELECT count(*) FROM users_segments WHERE user_id = $1", userID).Scan(&segmentsCount)

	segments := make([]string, 0, segmentsCount)
	for rows.Next() {
		segmentName := ""
		err := rows.Scan(&segmentName)
		if err != nil {
			return nil, fmt.Errorf("SQL: get segments: %w", err)
		}
		segments = append(segments, segmentName)
	}

	return segments, nil
}

func (r *Repo) CreateSegment(ctx context.Context, name string) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO segments (name) VALUES ($1)", name)
	if err != nil {
		return fmt.Errorf("SQL: create segment: %w", err)
	}
	return nil
}

func (r *Repo) AddUserSegments(ctx context.Context, addSegments []string, userID int) error {
	for _, segmentName := range addSegments {
		_, err := r.pool.Exec(ctx, "INSERT INTO users_segments (user_id, segment_name) VALUES ($1, $2)", userID, segmentName)
		if err != nil {
			return fmt.Errorf("SQL: add user segments: %w", err)
		}
	}
	return nil
}

func (r *Repo) RemoveUserSegments(ctx context.Context, addSegments []string, userID int) error {
	for _, segmentName := range addSegments {
		_, err := r.pool.Exec(ctx, "DELETE FROM users_segments WHERE user_id = $1 AND segment_name = $2", userID, segmentName)
		if err != nil {
			return fmt.Errorf("SQL: remove user segments: %w", err)
		}
	}
	return nil
}

func (r *Repo) DeleteSegment(ctx context.Context, name string) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM segments WHERE name = $1", name)
	if err != nil {
		return fmt.Errorf("SQL: delete segment: %w", err)
	}
	return nil
}
