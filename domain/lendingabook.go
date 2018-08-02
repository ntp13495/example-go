package domain

import (
	"time"
)

// LendingBooks decribes book in system
type LendingBooks struct {
	Model
	BookID UUID      `json:"book_id"`
	UserID UUID      `json:"user_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
