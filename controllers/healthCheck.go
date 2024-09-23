package controllers

import (
	"net/http"

	"jual-beli-motor/repository"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

var repo = repository.Database{}

func HealthCheck(ctx *gin.Context) {
	data := repo.GetHealthCheck(ctx)
	res := Response{
		Code:    http.StatusOK,
		Message: data,
	}

	ctx.JSON(res.Code, res)
}
