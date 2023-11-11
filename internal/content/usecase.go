package content

import (
	"context"
	"forum-authentication/internal/content/usecases"
)

type Cases struct {
	PostCases
	CommentCases
	CategoryCases
}

func NewCases(cr Repo) *Cases {
	return &Cases{
		PostCases:     usecases.NewPostCase(cr.PostRepo),
		CommentCases:  usecases.NewCommentCase(cr.CommentRepo),
		CategoryCases: usecases.NewCategoryCase(cr.CategoryRepo),
	}
}

type PostCases interface {
	Create(ctx context.Context) (*Post, error)
	Receive(ctx context.Context) (*Post, error)
	BulkReceive(ctx context.Context) ([]*Post, error)
	Update(ctx context.Context) (*Post, error)
}

type CommentCases interface {
	Create(ctx context.Context) (*Comment, error)
	BulkReceive(ctx context.Context) ([]*Comment, error)
	Update(ctx context.Context) (*Comment, error)
}

type CategoryCases interface {
	Receive(ctx context.Context) (*Comment, error)
	BulkReceive(ctx context.Context) ([]*Comment, error)
}
