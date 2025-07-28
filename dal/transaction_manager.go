package dal

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type TxFunc func(ctx context.Context, tx sqlx.ExtContext) error

type TransactionManager interface {
	WithTransaction(context.Context, TxFunc) error
}

type TransactionManagerImpl struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewTransactionManager(db *sqlx.DB, logger *zap.Logger) TransactionManager {
	return &TransactionManagerImpl{
		db:     db,
		logger: logger,
	}
}

func (m *TransactionManagerImpl) WithTransaction(ctx context.Context, txFunc TxFunc) error {
	tx, err := m.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("cannot start transaction: %w", err)
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
