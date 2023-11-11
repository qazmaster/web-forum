package content

import (
	"github.com/google/uuid"
	"time"
)

type CategoryID int

type Post struct {
	ID           int
	Title        string
	Body         string
	UserID       uuid.UUID
	CategoriesID []CategoryID
	CreatedTime  time.Time
	Votes        int
}

type Comment struct {
	ID          int
	Body        string
	UserID      uuid.UUID
	PostID      int
	CreatedTime time.Time
	Votes       int
}

type Category struct {
	ID   CategoryID
	Name string
}
