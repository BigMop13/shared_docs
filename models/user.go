package models

import "time"

type User struct {
	Username string
	Created  time.Time
}

func NewUser(username string, createdAt time.Time) *User {
	return &User{
		Username: username,
		Created:  createdAt,
	}
}
