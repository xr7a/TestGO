package config

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(NewConfig),
)

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
