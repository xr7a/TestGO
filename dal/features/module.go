package features

import (
	"awesomeProject/dal/features/departments"
	"awesomeProject/dal/features/passports"
	"awesomeProject/dal/features/users"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(users.NewWriteRepositoryImpl),
	fx.Provide(passports.NewWriteRepositoryImpl),
	fx.Provide(departments.NewWriteRepositoryImpl),
)
