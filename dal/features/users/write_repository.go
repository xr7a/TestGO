package users

import (
	"context"
	_ "embed"
	"github.com/jmoiron/sqlx"
)

type WriteRepository interface {
	CreateUser(context.Context, sqlx.ExtContext, DbUser) (DbUser, error)
}

type WriteRepositoryImpl struct{}

func NewWriteRepositoryImpl() WriteRepository {
	return &WriteRepositoryImpl{}
}

//go:embed write_sql/create_user.sql
var sqlCreateUser string

func (u *WriteRepositoryImpl) CreateUser(ctx context.Context, ext sqlx.ExtContext, user DbUser) (DbUser, error) {
	rows, err := sqlx.NamedQueryContext(ctx, ext, sqlCreateUser, user)
	if err != nil {
		return DbUser{}, err
	}
	defer rows.Close()

	if !rows.Next() {
		return DbUser{}, rows.Err()
	}

	err = rows.Scan(&user.Id)
	if err != nil {
		return DbUser{}, err
	}

	return user, nil
}
