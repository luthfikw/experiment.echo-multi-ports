package main

import (
	"go.uber.org/fx"

	"github.com/luthfikw/experiment.echo-multi-ports/internal/app/debugger"
	"github.com/luthfikw/experiment.echo-multi-ports/internal/app/httpprivate"
	"github.com/luthfikw/experiment.echo-multi-ports/internal/app/httppublic"
	"github.com/luthfikw/experiment.echo-multi-ports/internal/app/logger"
)

func main() {
	app := fx.New(
		logger.AddLoggerZap(),
		debugger.AddProfilerStatsviz(&debugger.ProfilerStatsvizArgs{
			Port: 4999,
		}),
		httpprivate.AddHTTPPrivate(&httpprivate.HTTPPrivateArgs{
			Port: 5001,
		}),
		httppublic.AddHTTPPublic(&httppublic.HTTPPublicArgs{
			Port: 5002,
		}),
	)
	app.Run()
}
