package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StarsStore struct {
	db *pgxpool.Pool
}

func NewStarsStore(db *pgxpool.Pool) *StarsStore {
	return &StarsStore{db: db}
}

func (s *StarsStore) Balance(ctx context.Context, userID int64) (int64, error) {
	var balance int64

	err := s.db.QueryRow(
		ctx,
		`SELECT balance FROM user_stars WHERE user_id = $1`,
		userID,
	).Scan(&balance)

	if err == pgx.ErrNoRows {
		return 0, nil
	}

	return balance, err
}
