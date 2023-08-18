package httppublic

import "github.com/labstack/echo/v4"

func (ox *httpPublicHandler) Welcome(ctx echo.Context) (err error) {
	ctx.JSON(200, "Welcome to public port")
	return
}
