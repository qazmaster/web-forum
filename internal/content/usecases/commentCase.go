package usecases

import (
	"context"
	"forum-authentication/internal/content"
)

type commentCase struct {
	repo *content.CommentRepo
}

func NewCommentCase(cr content.CommentRepo) content.CommentCases {
	return &commentCase{
		repo: &cr,
	}
}

func (c commentCase) Create(ctx context.Context) (*content.Comment, error) {
	//TODO implement me
	return nil, nil
}

func (c commentCase) BulkReceive(ctx context.Context) ([]*content.Comment, error) {
	//TODO implement me
	return nil, nil
}

func (c commentCase) Update(ctx context.Context) (*content.Comment, error) {
	//TODO implement me
	return nil, nil
}
