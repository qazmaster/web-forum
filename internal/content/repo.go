package content

import (
	"database/sql"
	"forum-authentication/internal/content/sqlite"
	"forum-authentication/internal/user"
	"github.com/google/uuid"
)

type Repo struct {
	PostRepo
	CommentRepo
	CategoryRepo
}

func NewSqlRepo(db *sql.DB) *Repo {
	return &Repo{
		PostRepo:     sqlite.NewPostRepo(db),
		CommentRepo:  sqlite.NewCommentRepo(db),
		CategoryRepo: sqlite.NewCategoryRepo(db),
	}
}

type PostRepo interface {
	Create(p *Post) (*Post, error)
	GetByID(i int) (*Post, error)
	GetByCategoryID(i int) ([]*Post, error)
	GetAll(userID uuid.UUID) ([]*Post, error)
	UpVote(v *user.Vote) error
	DownVote(v *user.Vote) error
}

type CommentRepo interface {
	Create(p *Comment) (*Comment, error)
	GetAllByPostID(i int) ([]*Comment, error)
	UpVote(v *user.Vote) error
	DownVote(v *user.Vote) error
}

type CategoryRepo interface {
	GetById(i int) (*Category, error)
	GetAll() ([]*Category, error)
	GetByPostID(i int) ([]*Category, error)
}
