package main

import (
	"awesomeProject/api"
	"awesomeProject/config"
	"awesomeProject/dal"
	"awesomeProject/services"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		config.Module,
		dal.Module,
		services.Module,
		api.Module,
	).Run()
}
