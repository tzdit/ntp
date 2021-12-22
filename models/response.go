package models

import (
	"net/http"
	"reflect"
	"time"
)

type ResponseDataModel struct {
	TimeStamp  time.Time `json:"timestamp"`
	StatusCode int32     `json:"status_code"`
	HttpStatus string    `json:"http_status"`
	Message    string    `json:"message"`
	Count      int32     `json:"count"`
	Data       interface {
	} `json:"data,omitempty"`
}

func ResponseSuccessData(data interface{}) *ResponseDataModel {
	val := reflect.ValueOf(data)
	msgDataResp := &ResponseDataModel{
		TimeStamp:  time.Now(),
		StatusCode: http.StatusOK,
		HttpStatus: http.StatusText(http.StatusOK),
		Message:    Success,
		Count:      int32(val.Len()),
		Data:       data,
	}
	return msgDataResp
}

func ResponseErrorData(data interface{}, errMsg string) *ResponseDataModel {
	val := reflect.ValueOf(data)
	msgDataResp := &ResponseDataModel{
		TimeStamp:  time.Now(),
		StatusCode: http.StatusInternalServerError,
		HttpStatus: http.StatusText(http.StatusInternalServerError),
		Message:    errMsg,
		Count:      int32(val.Len()),
		Data:       data,
	}
	return msgDataResp
}
