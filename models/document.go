package models

import "time"

type Document struct {
	Name     string
	Created  time.Time
	Modified time.Time
}

func NewDocument(name string) Document {
	return Document{
		Name: name,
	}
}
