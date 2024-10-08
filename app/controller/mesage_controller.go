package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopdeck.com/messaging_service/app/constant/response_status"
	"shopdeck.com/messaging_service/app/domain/dto/httprequest"
	"shopdeck.com/messaging_service/app/pkg"
	"shopdeck.com/messaging_service/app/service"
)

type MessageController interface {
	SendMessage(c *gin.Context)
	StartTest(c *gin.Context)
}

type MessageControllerImpl struct {
	messageService service.MessageService
}

func MessageControllerInit(messageService service.MessageService) MessageControllerImpl {
	return MessageControllerImpl{
		messageService: messageService,
	}
}

func (messageControllerImpl MessageControllerImpl) SendMessage(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request httprequest.PublishMessageRequest
	err := convertBodyFromJson(&request, c)
	if err != nil {
		// unmarshal error
		pkg.PanicException(response_status.InvalidRequest)
	}

	//request.Message = time.Now().String() + " " + request.Message
	err = messageControllerImpl.messageService.SendMessage(request)
	if err != nil {
		pkg.PanicException_("Error", err.Error())
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(response_status.Success, pkg.Null()))
}
