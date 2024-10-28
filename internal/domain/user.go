package domain

import "time"

type User struct {
	Id        string
	Username  string
	Password  string
	Friends   []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	User UserData `json:"user"`
}

type UserData struct {
	Username string `json:"username"`
}
