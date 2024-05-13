package router

import (
	"github.com/dMuto/umhelp-servidor-distopico/config"
	"github.com/dMuto/umhelp-servidor-distopico/server/controller"
	"github.com/labstack/echo/v4"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("")
	root.GET("/health", ctrl.HealthController.HealthCheck)
	root.POST("/create-account", ctrl.UserController.HandleNewUser)
}
