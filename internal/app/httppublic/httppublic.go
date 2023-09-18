package httppublic

import (
	"context"
	"fmt"
	"net"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HTTPPublicArgs struct {
	Port int
}

func AddHTTPPublic(args *HTTPPublicArgs) fx.Option {
	return fx.Options(
		fx.Supply(args),
		fx.Invoke(newHTTPPublic),
	)
}

type httpPublic struct {
	Logger   *zap.Logger
	Instance *echo.Echo
	Handler  *httpPublicHandler
}

func newHTTPPublic(lc fx.Lifecycle, args *HTTPPublicArgs, zlog *zap.Logger) {
	logger := zlog.With(
		zap.String("serverType", "public"),
		zap.String("port", fmt.Sprintf(":%d", args.Port)),
	)

	service := &httpPublic{
		Logger:   logger,
		Instance: echo.New(),
		Handler: &httpPublicHandler{
			Logger: logger,
		},
	}
	service.InitRoute()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			var listener net.Listener
			listener, err = net.Listen("tcp", fmt.Sprintf(":%d", args.Port))
			if err != nil {
				return
			}

			go func() {
				err := service.Instance.Server.Serve(listener)
				logger.Error("Failed to serve http public", zap.Error(err))
			}()

			return
		},
		OnStop: func(ctx context.Context) error {
			return service.Instance.Shutdown(ctx)
		},
	})
}
