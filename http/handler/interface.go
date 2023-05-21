package handler

import "github.com/ginkwok/ibook/service"

type HandlerStruct struct {
	svc service.Service
}

func NewHandler(s service.Service) *HandlerStruct {
	return &HandlerStruct{
		svc: s,
	}
}
