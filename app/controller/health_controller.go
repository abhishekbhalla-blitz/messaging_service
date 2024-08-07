package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopdeck.com/messaging_service/app/constant/response_status"
	"shopdeck.com/messaging_service/app/pkg"
)

type HealthController interface {
	PingPong(c *gin.Context)
}

type HealthControllerImpl struct {
}

func HealthControllerInit() *HealthControllerImpl {
	return &HealthControllerImpl{}
}

func (healthControllerImpl HealthControllerImpl) PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, pkg.BuildResponse(response_status.Success, pkg.Null()))
}
