package users

import "github.com/jmoiron/sqlx"

type WriteRepository interface {
	CreateUser(DbUser)
}

type WriteRepositoryImpl struct {
	db *sqlx.DB
}

func NewWriteRepositoryImpl(db *sqlx.DB) WriteRepository {
	return &WriteRepositoryImpl{db: db}
}

func (u *WriteRepositoryImpl) CreateUser(user DbUser) {

}
