package domain

import "time"

type Message struct {
	ID        int
	UserID    int64
	Role      string
	Content   string
	CreatedAt time.Time
	SelectedLanguage string
}