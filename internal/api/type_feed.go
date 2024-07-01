package api

import (
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	ErrorMsg  string    `json:"error"`
}

type CreateFeed struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}
