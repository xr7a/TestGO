package dal

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TxFunc func(ctx context.Context, tx sqlx.ExtContext) (interface{}, error)

type DatabaseManager interface {
	RunWithTransaction(context.Context, TxFunc) (interface{}, error)
	RunWithoutTransaction(context.Context, TxFunc) (interface{}, error)
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

func (m *DatabaseManagerImpl) RunWithTransaction(ctx context.Context, txFunc TxFunc) (interface{}, error) {
	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot start transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			m.logger.Panic("recovered from panic", zap.Any("panic", p))
			err := tx.Rollback()
			if err != nil {
				m.logger.Error("cannot rollback transaction after panic", zap.Error(err))
			}
			panic(p)
		} else if err != nil {
			m.logger.Error("error inside transaction operation", zap.Error(err))
		} else {
			commitErr := tx.Commit()
			if commitErr != nil {
				m.logger.Error("cannot commit transaction after panic", zap.Error(commitErr))
				err = fmt.Errorf("cannot commit transaction: %w", commitErr)
			}
		}
	}()

	return txFunc(ctx, tx)
}

func (m *DatabaseManagerImpl) RunWithoutTransaction(ctx context.Context, fn TxFunc) (interface{}, error) {
	return fn(ctx, m.db)
}

func RunQueryOperation[T any](ctx context.Context, dbManager DatabaseManager, fn TxFunc, withTransaction bool) (T, error) {
	uniFunc := func(ctx context.Context, exec sqlx.ExtContext) (interface{}, error) {
		result, err := fn(ctx, exec)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	var zero T
	var result interface{}
	var err error

	if withTransaction {
		result, err = dbManager.RunWithTransaction(ctx, uniFunc)
	} else {
		result, err = dbManager.RunWithoutTransaction(ctx, uniFunc)
	}
	if err != nil {
		return zero, err
	}

	typedResult, ok := result.(T)
	if !ok {
		return zero, fmt.Errorf("cannot cast result to %T", result, zero)
	}
	return typedResult, nil
}

func RunExecOperation(ctx context.Context, dbManager DatabaseManager, fn TxFunc, withTransaction bool) error {
	uniFunc := func(ctx context.Context, exec sqlx.ExtContext) (interface{}, error) {
		result, err := fn(ctx, exec)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	var err error
	if withTransaction {
		_, err = dbManager.RunWithTransaction(ctx, uniFunc)
	} else {
		_, err = dbManager.RunWithoutTransaction(ctx, uniFunc)
	}

	if err != nil {
		return err
	}
	return nil
}
