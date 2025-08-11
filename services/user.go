package services

import (
	"awesomeProject/config"
	"awesomeProject/dal"
	"awesomeProject/dal/features/passports"
	"awesomeProject/dal/features/users"
	"awesomeProject/mappers"
	"awesomeProject/models"
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.UserPassport) (models.SuccessCreateUser, error)
}

type UserServiceImpl struct {
	dbManager               dal.DatabaseManager
	userWriteRepository     users.WriteRepository
	passportWriteRepository passports.WriteRepository
	logger                  *zap.Logger
}

func NewUserServiceImpl(dbManager dal.DatabaseManager, userWriteRepository users.WriteRepository, passportWriteRepository passports.WriteRepository, logger *zap.Logger, config config.Config) UserService {
	return &UserServiceImpl{
		dbManager:               dbManager,
		userWriteRepository:     userWriteRepository,
		passportWriteRepository: passportWriteRepository,
		logger:                  logger,
	}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, userPassport models.UserPassport) (models.SuccessCreateUser, error) {
	return dal.RunQueryOperation[models.SuccessCreateUser](ctx, u.dbManager, func(ctx context.Context, tx sqlx.ExtContext) (models.SuccessCreateUser, error) {
		dbUser := mappers.MapToDbUser(userPassport)
		createdUser, err := u.userWriteRepository.CreateUser(ctx, tx, dbUser)
		if err != nil {
			u.logger.Error("Failed to create user", zap.Error(err))
			return models.SuccessCreateUser{}, err
		}

		user := mappers.MapToDomainSuccessCreateUser(createdUser)

		passport, err := u.passportWriteRepository.CreatePassport(ctx, tx, mappers.MapToDbPassport(userPassport, user.Id))
		if err != nil {
			u.logger.Error("Failed to create passport", zap.Error(err))
			return models.SuccessCreateUser{}, err
		}

		user.Passport = mappers.MapToDomainPassport(passport)

		return user, nil
	}, true)
}
