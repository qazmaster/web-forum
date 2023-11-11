package sqlite

import (
	"database/sql"
	"forum-authentication/internal/content"
	"forum-authentication/internal/user"
	"github.com/google/uuid"
)

type postRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) content.PostRepo {
	return &postRepo{
		db: db,
	}
}

func (pr postRepo) Create(p *content.Post) (*content.Post, error) {
	//TODO implement me
	return nil, nil
}
func (pr postRepo) GetByID(i int) (*content.Post, error) {
	//TODO implement me
	return nil, nil
}
func (pr postRepo) GetByCategoryID(i int) ([]*content.Post, error) {
	//TODO implement me
	return nil, nil
}
func (pr postRepo) GetAll(userID uuid.UUID) ([]*content.Post, error) {
	//TODO implement me
	return nil, nil
}
func (pr postRepo) UpVote(v *user.Vote) error {
	//TODO implement me
	return nil
}
func (pr postRepo) DownVote(v *user.Vote) error {
	//TODO implement me
	return nil
}
