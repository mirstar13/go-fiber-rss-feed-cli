package api

import (
	"time"
)

type User struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
	ErrorMsg  string    `json:"error"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
