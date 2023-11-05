package echo

import (
	"github.com/Raajheer1/hackutd-x/m/v2/pkg/echo/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func OnlineCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.MessageResponse{
		Message: "I am online!",
	})
}
