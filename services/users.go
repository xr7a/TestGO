package services

import (
	"awesomeProject/config"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserService interface {
	GetConfig()
}

type UserServiceImpl struct {
	db     *sqlx.DB
	logger *zap.Logger
	config *config.Config
}

func NewUserServiceImpl(db *sqlx.DB, logger *zap.Logger, config config.Config) UserService {
	return &UserServiceImpl{
		db:     db,
		logger: logger,
		config: &config,
	}
}

func (u *UserServiceImpl) GetConfig() {
	u.logger.Info("Говно", zap.String("Connection", u.config.ConnectionString))
}
