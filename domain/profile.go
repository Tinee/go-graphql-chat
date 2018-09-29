package domain

import (
	"time"
)

type Profile struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
}
