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
func CreateErrorResponse(err iterror.ErrorException, data ...interface{}) model.ServiceResponse {
	var errMsg interface{}
	if len(data) > 0 {
		errMsg = data[0]
	} else {
		errMsg = err.Error()
	}
	return model.ServiceResponse{
		Status:       false,
		ServiceCode:  serviceCode,
		Data:         nil,
		ErrorCode:    err.GetCode(),
		ErrorName:    err.GetName(),
		ErrorMessage: errMsg,
	}
}
