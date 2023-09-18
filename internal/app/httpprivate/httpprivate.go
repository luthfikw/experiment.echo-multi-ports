package httpprivate

import (
	"context"
	"fmt"
	"net"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HTTPPrivateArgs struct {
	Port int
}

func AddHTTPPrivate(args *HTTPPrivateArgs) fx.Option {
	return fx.Options(
		fx.Supply(args),
		fx.Invoke(newHTTPPrivate),
	)
}

type httpPrivate struct {
	Logger   *zap.Logger
	Instance *echo.Echo
	Handler  *httpPrivateHandler
}

func newHTTPPrivate(lc fx.Lifecycle, args *HTTPPrivateArgs, zlog *zap.Logger) {
	logger := zlog.With(
		zap.String("serverType", "private"),
		zap.String("port", fmt.Sprintf(":%d", args.Port)),
	)

	service := &httpPrivate{
		Logger:   logger,
		Instance: echo.New(),
		Handler: &httpPrivateHandler{
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
				logger.Error("Failed to serve http private", zap.Error(err))
			}()

			return
		},
		OnStop: func(ctx context.Context) error {
			return service.Instance.Shutdown(ctx)
		},
	})
}
