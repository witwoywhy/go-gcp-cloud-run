package handler

import (
	"go-gcp-cloud-run/services/greeting"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/witwoywhy/go-cores/gins"
)

type greetingHandler struct {
	service greeting.Service
}

func NewGreetingHandler(service greeting.Service) *greetingHandler {
	return &greetingHandler{
		service: service,
	}
}

func (h *greetingHandler) Handle(ctx *gin.Context) {
	l := gins.NewLogFromCtx(ctx)
	var request greeting.Request
	request.Name = ctx.Query("name")

	response, err := h.service.Execute(request, l)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
