package usecases

import (
	"context"
	"forum-authentication/internal/content"
)

type categoryCase struct {
	repo *content.CategoryRepo
}

func NewCategoryCase(cr content.CategoryRepo) content.CategoryCases {
	return &categoryCase{
		repo: cr,
	}
}

func (c categoryCase) Receive(ctx context.Context) (*content.Comment, error) {
	//TODO implement me
	return nil, nil
}

func (c categoryCase) BulkReceive(ctx context.Context) ([]*content.Comment, error) {
	//TODO implement me
	return nil, nil
}
