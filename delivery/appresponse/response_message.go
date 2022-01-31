package appresponse

import (
	"log"
	"net/http"
)

type ResponseMessage struct {
	Status      string      `json:"status"`
	Description string      `json:"message"`
	Data        interface{} `json:"data"`
}

type ErrorMessage struct {
	ErrCode          int    `json:"code"`
	ErrorDescription string `json:"message"`
}

func NewResponseMessage(status string, description string, data interface{}) *ResponseMessage {
	return &ResponseMessage{
		Status:      status,
		Description: description,
		Data:        data,
	}
}

func NewUnauthorizedError(err error, message string) *ErrorMessage {
	em := &ErrorMessage{
		ErrCode:          http.StatusUnauthorized,
		ErrorDescription: message,
	}
	log.Println(err.Error())
	return em
}

func NewInternalServerError(err error, message string) *ErrorMessage {
	em := &ErrorMessage{
		ErrCode:          http.StatusInternalServerError,
		ErrorDescription: message,
	}
	log.Println(err.Error())
	return em
}

func NewBadRequestError(err error, message string) *ErrorMessage {
	em := &ErrorMessage{
		ErrCode:          http.StatusBadRequest,
		ErrorDescription: message,
	}
	log.Println(err.Error())
	return em
}
