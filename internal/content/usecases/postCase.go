package usecases

import (
	"context"
	"forum-authentication/internal/content"
)

type postCase struct {
	repo *content.PostRepo
}

func NewPostCase(pr content.PostRepo) content.PostCases {
	return &postCase{
		repo: &pr,
	}
}

func (p postCase) Create(ctx context.Context) (*content.Post, error) {
	//TODO implement me
	return nil, nil
}

func (p postCase) Receive(ctx context.Context) (*content.Post, error) {
	//TODO implement me
	return nil, nil
}

func (p postCase) BulkReceive(ctx context.Context) ([]*content.Post, error) {
	//TODO implement me
	return nil, nil
}

func (p postCase) Update(ctx context.Context) (*content.Post, error) {
	//TODO implement me
	return nil, nil
}
