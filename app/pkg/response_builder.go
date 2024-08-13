package pkg

import (
	"shopdeck.com/messaging_service/app/constant/response_status"
	"shopdeck.com/messaging_service/app/domain/dto/httprequest"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus response_status.ResponseStatus, data T) httprequest.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) httprequest.ApiResponse[T] {
	return httprequest.ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
