package model

import "time"

type thread struct {
	id      string    `json:"id"`
	title   string    `json:"title"` //optional
	content string    `json:"name"`
	created time.Time `json:"name"`
}
