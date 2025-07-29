package api

import (
	"awesomeProject/api/handlers"
	"awesomeProject/services"
	"context"
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

var Module = fx.Options(
	fx.Provide(MainRouter),
	fx.Invoke(StartServer))

func MainRouter(userService services.UserService, departmentService services.DepartmentService) *mux.Router {
	router := mux.NewRouter()

	departmentRouter := router.PathPrefix("/department").Subrouter()
	handlers.HandleDepartmentRoutes(departmentRouter, departmentService)

	usersRouter := router.PathPrefix("/users").Subrouter()
	handlers.HandleUserRoutes(usersRouter, userService)

	pingRouter := router.PathPrefix("/ping").Subrouter()
	handlers.HandlePingRoutes(pingRouter)

	return router
}

func StartServer(lc fx.Lifecycle, logger *zap.Logger, router *mux.Router) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("starting server", zap.String("addr", server.Addr))

			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error(err.Error())
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("stopping server", zap.String("addr", server.Addr))

			err := server.Shutdown(ctx)
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			return nil
		},
	})
}
