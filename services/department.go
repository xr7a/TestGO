package services

import (
	"awesomeProject/dal"
	"awesomeProject/dal/features/departments"
	"awesomeProject/mappers"
	"awesomeProject/models"
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DepartmentService interface {
	CreateDepartment(context.Context, models.Department) error
}

type DepartmentServiceImpl struct {
	dbManager                 dal.DatabaseManager
	departmentWriteRepository departments.WriteRepository
	logger                    *zap.Logger
}

func NewDepartmentServiceImpl(dbManager dal.DatabaseManager, departmentWriteRepository departments.WriteRepository, logger *zap.Logger) DepartmentService {
	return &DepartmentServiceImpl{
		dbManager:                 dbManager,
		departmentWriteRepository: departmentWriteRepository,
		logger:                    logger,
	}
}

func (s *DepartmentServiceImpl) CreateDepartment(ctx context.Context, department models.Department) error {
	return dal.RunExecOperation(ctx, s.dbManager, func(ctx context.Context, tx sqlx.ExtContext) error {
		return s.departmentWriteRepository.CreateDepartment(ctx, tx, mappers.MapToDbDepartment(department))
	}, false)
}
