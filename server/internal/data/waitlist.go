package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Waitlist struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type WaitlistModel struct {
	DB *sql.DB
}

func (m WaitlistModel) Insert(waitlist *Waitlist) error {
	query := `
		INSERT INTO waitlist (email)
		VALUES ($1)
		RETURNING id, created_at
	`

	args := []any{waitlist.Email}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&waitlist.ID, &waitlist.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "waitlist_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}
