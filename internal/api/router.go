package api

import (
	"github.com/ivofreitas/device-api/internal/api/device"
	"github.com/ivofreitas/device-api/internal/api/middleware"
	"github.com/ivofreitas/device-api/internal/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

func register(echo *echo.Echo) {
	deviceGroup(echo)
}

func deviceGroup(echo *echo.Echo) {
	deviceServ := device.NewService(device.NewRepository())
	createDevice := middleware.NewHandler(deviceServ.Create, http.StatusCreated, domain.CreateDTO{})

	group := echo.Group("/device")
	group.POST("/", createDevice.Handle)
}
