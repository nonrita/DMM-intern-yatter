package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

func (s *status) Create(ctx context.Context, tx *sqlx.Tx, sta *object.Status) error {
	_, err := s.db.Exec("insert into status (account_id, content, url) values (?, ?, ?)",sta.AccountID, sta.Content, sta.URL)
	if err != nil {
		return fmt.Errorf("failed to insert account: %w", err)
	}

	return nil
}