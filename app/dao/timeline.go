package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	timeline struct {
		db *sqlx.DB
	}
)

var _ repository.Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

func (t *timeline) FindAll(ctx context.Context, isMedia bool, maxID int, sinceID int, limit int) ([]object.TimelineItem, error) {
	var entities []object.TimelineItem
	query := `select s.id, a.id as "account.id", a.username as "account.username", a.create_at as "account.create_at", s.url, s.content, s.created_at from status as s inner join account as a on a.id=s.account_id where ? < s.id and s.id < ? `
	if isMedia {
		query = query + ` and s.url != "" `
	}
	query = query + ` limit ? `
	err := t.db.SelectContext(ctx, &entities, query, sinceID, maxID, limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entities, nil
}