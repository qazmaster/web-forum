package sqlite

import (
	"database/sql"
	"forum-authentication/internal/content"
)

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) content.CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c categoryRepo) GetById(i int) (*content.Category, error) {
	//TODO implement me
	return nil, nil
}

func (c categoryRepo) GetAll() ([]*content.Category, error) {
	//TODO implement me
	return nil, nil
}

func (c categoryRepo) GetByPostID(i int) ([]*content.Category, error) {
	//TODO implement me
	return nil, nil
}
