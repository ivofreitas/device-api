package api

import (
	"github.com/ivofreitas/device-api/config/db"
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
	deviceServ := device.NewService(device.NewRepository(db.NewPostgresConnection()))
	createHdl := middleware.NewHandler(deviceServ.Create, http.StatusCreated, domain.Device{})
	updateHdl := middleware.NewHandler(deviceServ.Update, http.StatusOK, domain.Update{})
	patchHdl := middleware.NewHandler(deviceServ.Patch, http.StatusOK, domain.Patch{})
	getAllHdl := middleware.NewHandler(deviceServ.GetAll, http.StatusOK, nil)
	getByIdHdl := middleware.NewHandler(deviceServ.GetById, http.StatusOK, domain.GetById{})
	getByBrandHdl := middleware.NewHandler(deviceServ.GetByBrand, http.StatusOK, domain.GetByBrand{})
	getByStateHdl := middleware.NewHandler(deviceServ.GetByState, http.StatusOK, domain.GetByState{})
	deleteHdl := middleware.NewHandler(deviceServ.Delete, http.StatusOK, domain.Delete{})

	group := echo.Group("/devices")
	group.POST("", createHdl.Handle)
	group.PUT("/:id", updateHdl.Handle)
	group.PATCH("/:id", patchHdl.Handle)
	group.GET("", getAllHdl.Handle)
	group.GET("/:id", getByIdHdl.Handle)
	group.GET("/:brand", getByBrandHdl.Handle)
	group.GET("/:state", getByStateHdl.Handle)
	group.DELETE("/:id", deleteHdl.Handle)
}
