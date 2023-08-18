package httpprivate

import "github.com/labstack/echo/v4"

func (ox *httpPrivateHandler) Welcome(ctx echo.Context) (err error) {
	ctx.JSON(200, "Welcome to private port")
	return
}
