package controller

import (
	model "m9-backstore-service/models"

	iterror "github.com/M9nood/go-iterror"
)

var serviceCode = "00"

func CreateSuccessResponse(data interface{}) model.ServiceResponse {
	return model.ServiceResponse{
		Status:      true,
		StatusCode:  "200",
		ServiceCode: serviceCode,
		Data:        data,
	}
}

// CreateErrorResponse - create fail response
func CreateErrorResponse(err iterror.ErrorException) model.ServiceResponse {
	return model.ServiceResponse{
		Status:       false,
		ServiceCode:  serviceCode,
		Data:         nil,
		ErrorCode:    err.GetCode(),
		ErrorName:    err.GetName(),
		ErrorMessage: err.Error(),
	}
}
