package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)


type Status interface {
	Create(ctx context.Context, accountID int, content string, url *string) (*CreateStatusDTO, error)
}

type status struct {
	db          *sqlx.DB
	statusRepo repository.Status
}


type CreateStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:          db,
		statusRepo: statusRepo,
	}
}

func (s *status) Create(ctx context.Context,accountID int, content string, url *string) (*CreateStatusDTO, error) {
	sta := object.NewStatus(content)
	sta.AccountID = accountID
	sta.URL = url

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func ()  {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	if err := s.statusRepo.Create(ctx, tx, sta); err != nil {
		return nil, err
	}

	return &CreateStatusDTO{
		Status: sta,
	}, nil
}