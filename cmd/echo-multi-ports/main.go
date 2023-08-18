package main

import (
	"go.uber.org/fx"

	"github.com/luthfikw/experiment.echo-multi-ports/internal/app/httpprivate"
	"github.com/luthfikw/experiment.echo-multi-ports/internal/app/httppublic"
)

func main() {
	app := fx.New(
		httpprivate.AddHTTPPrivate(&httpprivate.HTTPPrivateArgs{
			Port: 5001,
		}),
		httppublic.AddHTTPPublic(&httppublic.HTTPPublicArgs{
			Port: 5002,
		}),
	)
	app.Run()
}
