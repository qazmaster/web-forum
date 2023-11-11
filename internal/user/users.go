package user

import (
	"github.com/google/uuid"
	"time"
)

type VoteType string

type User struct {
	ID             uuid.UUID
	Username       string
	HashedPassword []byte
	Email          string
	CreatedTime    time.Time
}

type Session struct {
	UserID      uuid.UUID
	Token       string
	CreatedTime time.Time
	Expires     time.Time
}

type Vote struct {
	ID        int
	UserID    uuid.UUID
	Type      VoteType
	PostID    int
	CommentID int
}
