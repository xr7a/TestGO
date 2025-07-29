package departments

import (
	"context"
	_ "embed"
	"github.com/jmoiron/sqlx"
)

type WriteRepository interface {
	CreateDepartment(context.Context, sqlx.ExtContext, DbDepartment) error
}

type WriteRepositoryImpl struct{}

func NewWriteRepositoryImpl() WriteRepository {
	return &WriteRepositoryImpl{}
}

//go:embed write_sql/create_department.sql
var sqlCreateDepartment string

func (r *WriteRepositoryImpl) CreateDepartment(ctx context.Context, ext sqlx.ExtContext, department DbDepartment) error {
	_, err := sqlx.NamedExecContext(ctx, ext, sqlCreateDepartment, department)
	return err
}
