package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func AddLoggerZap() fx.Option {
	return fx.Options(
		fx.Provide(newLoggerZap),
	)
}

func newLoggerZap() (logger *zap.Logger, err error) {
	logger = zap.L()
	return
}
