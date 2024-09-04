package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Status interface{
	FindByID(ctx context.Context, id int) (*object.Status, error)
	Create(ctx context.Context, tx *sqlx.Tx, sta *object.Status) error
}