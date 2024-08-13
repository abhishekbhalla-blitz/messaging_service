package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopdeck.com/messaging_service/app/constant/response_status"
	"shopdeck.com/messaging_service/app/domain/dto/httprequest"
	"shopdeck.com/messaging_service/app/pkg"
	"shopdeck.com/messaging_service/app/service"
	"shopdeck.com/messaging_service/config"
)

type HealthController interface {
	ServiceHealthCheck(c *gin.Context)
}

type HealthControllerImpl struct {
	messageService service.MessageService
}

func HealthControllerInit(messageService service.MessageService) *HealthControllerImpl {
	return &HealthControllerImpl{
		messageService: messageService,
	}
}

func (healthController HealthControllerImpl) ServiceHealthCheck(c *gin.Context) {
	configuration := config.GetConfiguration()

	heathCheckResponseBuilder := httprequest.NewHealthCheckResponseBuilder()
	heathCheckResponseBuilder.SetServerHealthStatus(true)

	if configuration.Producer.Primary.Enabled {
		isHealthy := healthController.messageService.GetPrimaryProducerHealth()
		heathCheckResponseBuilder.SetPrimaryHealthStatus(true, isHealthy)
	}

	if configuration.Producer.Fallback.Enabled {
		isHealthy := healthController.messageService.GetFallbackProducerHealth()
		heathCheckResponseBuilder.SetFallbackHealthStatus(true, isHealthy)

		circuitBreakerState := healthController.messageService.GetCircuitBreakerState()
		heathCheckResponseBuilder.SetCircuitBreakerState(circuitBreakerState)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(response_status.Success, heathCheckResponseBuilder.Build()))
}
