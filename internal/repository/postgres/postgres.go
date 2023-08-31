package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgconn"
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
	rows, err := r.pool.Query(ctx, "SELECT segments.name FROM users_segments join segments on segment_id = segments.id WHERE user_id = $1", userID)
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return errors.New("segment already exists")
			}
		}
		return fmt.Errorf("SQL: create segment: %w", err)
	}
	return nil
}

func (r *Repo) AddUserSegments(ctx context.Context, addSegments []string, userID int) ([]string, error) { // TODO: decision
	query := "INSERT INTO users_segments (user_id, segment_id) VALUES ($1, (SELECT id FROM segments WHERE name = $2));"
	notAddedSegments := make([]string, 0, 5)

	for _, segmentName := range addSegments {
		_, err := r.pool.Exec(ctx, query, userID, segmentName)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				log.Println("as")
				notAddedSegments = append(notAddedSegments, segmentName)
			} else {
				log.Println("not as")
			}
		}
	}

	log.Println("notAddedSegments", notAddedSegments)
	if len(notAddedSegments) > 0 {
		return notAddedSegments, errors.New("SQL: add user segments: non existing segments")
	}

	return nil, nil
}

func (r *Repo) RemoveUserSegments(ctx context.Context, removeSegments []string, userID int) ([]string, error) { // TODO: decision
	query := "DELETE FROM users_segments WHERE user_id = $1 and segment_id = (SELECT id FROM segments WHERE name = $2);"
	notDeletedSegments := make([]string, 0, 5)

	for _, segmentName := range removeSegments {
		tag, _ := r.pool.Exec(ctx, query, userID, segmentName)
		if tag.RowsAffected() == 0 {
			notDeletedSegments = append(notDeletedSegments, segmentName)
		}
	}

	if len(notDeletedSegments) > 0 { // TODO: return to service; let service check length
		return notDeletedSegments, errors.New("SQL: remove user segments: not deleted segments")
	}
	return nil, nil
}

func (r *Repo) DeleteSegment(ctx context.Context, name string) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM segments WHERE name = $1", name)
	if err != nil {
		return fmt.Errorf("SQL: delete segment: %w", err)
	}
	return nil
}

func (r *Repo) SetPercentageSegments(ctx context.Context, name string, percentage int) ([]int, error) {
	query := `	INSERT INTO users_segments (user_id, segment_id)
					SELECT 	user_id,
  							(SELECT id FROM segments WHERE name = $1)
					FROM
						(WITH users AS
							(SELECT DISTINCT user_id FROM users_segments) 
						SELECT users.user_id
						FROM users
						ORDER BY random()
						LIMIT (SELECT count(*) FROM users)*0.01*$2) AS needed_users
					ORDER BY user_id
					ON CONFLICT DO NOTHING
					RETURNING user_id;`

	rows, err := r.pool.Query(ctx, query, name, percentage)
	if err != nil {
		return nil, fmt.Errorf("SQL: set percentage segments: %w", err)
	}

	var usersCount int
	r.pool.QueryRow(ctx, "select floor(count(*)*0.01*$1)::integer from (select distinct user_id from users_segments) as u;", percentage).Scan(&usersCount)

	affectedUsers := make([]int, 0, usersCount)
	for rows.Next() {
		var userID int
		err := rows.Scan(&userID)
		if err != nil {
			return nil, fmt.Errorf("SQL: SetPercentageSegments: get affected user: %w", err)
		}
		affectedUsers = append(affectedUsers, userID)
	}

	return affectedUsers, nil
}
