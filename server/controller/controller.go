package controller

import (
	"github.com/dMuto/umhelp-servidor-distopico/server/controller/health"
	"github.com/dMuto/umhelp-servidor-distopico/server/controller/user"
	"github.com/dMuto/umhelp-servidor-distopico/service"
	"github.com/dMuto/umhelp-servidor-distopico/util/resutil"
	"github.com/rs/zerolog"
)

type Controller struct {
	HealthController *health.Controller
	UserController   *user.Controller
}

func New(svc *service.Service, logger *zerolog.Logger) *Controller {
	resutil := resutil.New(logger)

	return &Controller{
		HealthController: health.New(resutil),
		UserController:   user.New(svc, resutil),
	}
}
