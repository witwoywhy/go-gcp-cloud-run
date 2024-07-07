package handler

import (
	"go-gcp-cloud-run/services/login"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/witwoywhy/go-cores/gins"
)

type loginHandler struct {
	service login.Service
}

func NewLoginHandler(service login.Service) *loginHandler {
	return &loginHandler{
		service: service,
	}
}

func (h *loginHandler) Handle(ctx *gin.Context) {
	l := gins.NewLogFromCtx(ctx)
	var request login.Request
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Error(err)
		return
	}

	response, err := h.service.Execute(request, l)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
