package object

import "time"

type TimelineItem struct {
	ID        int       `json:"id,omitempty"`
	Account 	Account		`db:"account"`
	URL       *string   `json:"url,omitempty" db:"url"`
	Content   string    `json:"status"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}