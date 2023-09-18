package debugger

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/arl/statsviz"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ProfilerStatsvizArgs struct {
	Port int
}

func AddProfilerStatsviz(args *ProfilerStatsvizArgs) fx.Option {
	return fx.Options(
		fx.Supply(args),
		fx.Invoke(newProfilerStatsviz),
	)
}

func newProfilerStatsviz(lc fx.Lifecycle, args *ProfilerStatsvizArgs, zlog *zap.Logger) (err error) {
	logger := zap.L().With(
		zap.String("serverType", "profiler"),
		zap.String("port", fmt.Sprintf(":%d", args.Port)),
	)

	mux := http.NewServeMux()
	if err = statsviz.Register(mux); err != nil {
		return
	}

	server := http.Server{
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			var listener net.Listener
			listener, err = net.Listen("tcp", fmt.Sprintf(":%d", args.Port))
			if err != nil {
				return
			}

			go func() {
				err := server.Serve(listener)
				logger.Error("Failed to serve profiler.statsviz server", zap.Error(err))
			}()

			return
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return
}
