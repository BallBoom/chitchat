package moudels

import "time"

type Session struct {
	ID        int
	UUID      string
	UserID    int
	Email     string
	CreatedAt time.Time
}
