package sqlite

import (
	"database/sql"
	"forum-authentication/internal/content"
	"forum-authentication/internal/user"
)

type commentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) content.CommentRepo {
	return &commentRepo{
		db: db,
	}
}
func (c commentRepo) Create(p *content.Comment) (*content.Comment, error) {
	//TODO implement me
	return nil, nil
}

func (c commentRepo) GetAllByPostID(i int) ([]*content.Comment, error) {
	//TODO implement me
	return nil, nil
}

func (c commentRepo) UpVote(v *user.Vote) error {
	//TODO implement me
	return nil
}

func (c commentRepo) DownVote(v *user.Vote) error {
	//TODO implement me
	return nil
}
