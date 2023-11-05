package ipad

import "github.com/labstack/echo/v4"

func Routes(r *echo.Group) {
	r.GET("/configs", getConfigs)
}
