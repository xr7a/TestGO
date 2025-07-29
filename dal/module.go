package dal

import (
	"awesomeProject/dal/features"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewDb),
	fx.Provide(NewTransactionManager),
	fx.Options(features.Module),
)
