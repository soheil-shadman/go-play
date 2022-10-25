package api

import (
	"myapp/config"

	"github.com/labstack/echo/v4"
)

func SendError(c echo.Context, code int, err string) error {
	return c.JSON(code, HttpAPIResponse{Code: code, Error: &err})
}
func SendResponse(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, HttpAPIResponse{Code: code, Data: data})
}
func SendAccessDenied(c echo.Context) error {
	return SendError(c, 400, "access denied")
}
func TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("api-token") != config.API_TOKEN {
			return SendError(c, 401, "access denied")
		}
		return next(c)
	}
}
func HasSystemToken(c echo.Context) bool {
	return (c.Request().Header.Get("system-token") == config.SYSTEM_TOKEN)
}
func HttpErrorHandler(err error, c echo.Context) {
	// c.Logger().Error(echo.HTTPError{Code: 500, Message: err.Error()})
	if he, ok := err.(*echo.HTTPError); ok {
		SendError(c, he.Code, he.Message.(string))
		return
	}
	SendError(c, 500, err.Error())
}
