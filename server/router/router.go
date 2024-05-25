package router

import (
	"github.com/dMuto/umhelp-servidor-distopico/config"
	"github.com/dMuto/umhelp-servidor-distopico/server/controller"
	"github.com/labstack/echo/v4"
)

func Register(cfg *config.Config, svr *echo.Echo, ctrl *controller.Controller) {
	root := svr.Group("/um-help-distopic")
	root.GET("/health", ctrl.HealthController.HealthCheck)
	
	userGroup := root.Group("/user")
	userGroup.POST("/create-account", ctrl.UserController.HandleNewUser)

	//walletGroup := root.Group("/wallet")
	//walletGroup.POST("/send-money",)
}
