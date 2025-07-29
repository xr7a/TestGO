package passports

import (
	"context"
	_ "embed"
	"github.com/jmoiron/sqlx"
)

type WriteRepository interface {
	CreatePassport(context.Context, sqlx.ExtContext, DbPassport) (DbPassport, error)
}

type WriteRepositoryImpl struct{}

func NewWriteRepositoryImpl() WriteRepository {
	return &WriteRepositoryImpl{}
}

//go:embed write_sql/create_passport.sql
var sqlCreatePassport string

func (p *WriteRepositoryImpl) CreatePassport(ctx context.Context, ext sqlx.ExtContext, passport DbPassport) (DbPassport, error) {
	row, err := sqlx.NamedQueryContext(ctx, ext, sqlCreatePassport, passport)
	if err != nil {
		return DbPassport{}, err
	}

	defer row.Close()
	if !row.Next() {
		return DbPassport{}, row.Err()
	}

	var result DbPassport
	if err := row.StructScan(&result); err != nil {
		return DbPassport{}, err
	}

	return result, nil
}
