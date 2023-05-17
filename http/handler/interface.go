package handler

import "github.com/ginkwok/ibook/service"

type handlerStruct struct {
	svc service.Service
}

func NewHandler(s service.Service) *handlerStruct {
	return &handlerStruct{
		svc: s,
	}
}
