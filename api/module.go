package api

import (
	"awesomeProject/api/handlers"
	"context"
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

var Module = fx.Options(
	fx.Provide(PingRouter),
	fx.Invoke(StartServer))

func PingRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)
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
