package dal

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewDb),
	fx.Provide(NewTransactionManager),
)
