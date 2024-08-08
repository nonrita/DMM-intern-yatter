package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface{
	FindAll(ctx context.Context, isMedia bool, maxID int, sinceID int, limit int) ([]object.TimelineItem, error)
}