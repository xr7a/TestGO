package dal

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DatabaseManager interface {
	RunWithTransaction(context.Context, func(context.Context, sqlx.ExtContext) error) error
	RunWithoutTransaction(context.Context, func(context.Context, sqlx.ExtContext) error) error
}

type DatabaseManagerImpl struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewTransactionManager(db *sqlx.DB, logger *zap.Logger) DatabaseManager {
	return &DatabaseManagerImpl{
		db:     db,
		logger: logger,
	}
}

func (m *DatabaseManagerImpl) RunWithTransaction(ctx context.Context, txFunc func(ctx context.Context, tx sqlx.ExtContext) error) error {
	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("cannot start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			m.logger.Panic("recovered from panic", zap.Any("panic", p))
			rerr := tx.Rollback()
			if rerr != nil {
				m.logger.Error("cannot rollback transaction after panic", zap.Error(rerr))
			}
			panic(p)
		} else if err != nil {
			m.logger.Error("error inside transaction operation", zap.Error(err))
			rerr := tx.Rollback()
			if rerr != nil {
				m.logger.Error("cannot rollback transaction", zap.Error(rerr))
			}
		} else {
			commitErr := tx.Commit()
			if commitErr != nil {
				m.logger.Error("cannot commit transaction after panic", zap.Error(commitErr))
				err = fmt.Errorf("cannot commit transaction: %w", commitErr)
			}
		}
	}()

	err = txFunc(ctx, tx)
	return err
}

func (m *DatabaseManagerImpl) RunWithoutTransaction(ctx context.Context, fn func(ctx context.Context, exec sqlx.ExtContext) error) error {
	return fn(ctx, m.db)
}

func RunQueryOperation[T any](ctx context.Context, dbManager DatabaseManager, fn func(ctx context.Context, exec sqlx.ExtContext) (T, error), withTransaction bool) (T, error) {
	var (
		result T
		err    error
	)

	callback := func(ctx context.Context, exec sqlx.ExtContext) error {
		result, err = fn(ctx, exec)
		return err
	}

	if withTransaction {
		err = dbManager.RunWithTransaction(ctx, callback)
	} else {
		err = dbManager.RunWithoutTransaction(ctx, callback)
	}

	if err != nil {
		var zero T
		return zero, err
	}
	return result, nil
}

func RunExecOperation(ctx context.Context, dbManager DatabaseManager, fn func(ctx context.Context, exec sqlx.ExtContext) error, withTransaction bool) error {
	if withTransaction {
		return dbManager.RunWithTransaction(ctx, fn)
	}
	return dbManager.RunWithoutTransaction(ctx, fn)
}
