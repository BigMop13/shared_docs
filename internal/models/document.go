package models

import "time"

type Document struct {
	Content   string    `json:"content"`
	Modified  time.Time `json:"modified"`
	LastSaved time.Time `json:"lastSaved"`
}
