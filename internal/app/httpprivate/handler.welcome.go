package httpprivate

import "github.com/labstack/echo/v4"

func (ox *httpPrivateHandler) Welcome(ctx echo.Context) (err error) {
	ctx.JSON(200, "Welcome to private port")
	return
}

func (ox *httpPrivateHandler) WelcomeA(ctx echo.Context) (err error) {
	ctx.JSON(200, "Welcome to private port with handler A")
	return
}

func (ox *httpPrivateHandler) WelcomeB(ctx echo.Context) (err error) {
	ctx.JSON(200, "Welcome to private port with handler B")
	return
}
