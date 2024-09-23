package controllers

import (
	"fmt"
	"jual-beli-motor/models"
	"jual-beli-motor/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBike(ctx *gin.Context) {
	res := Response{}
	payload := models.ReqBike{}

	if err := ctx.ShouldBind(&payload); err != nil {
		fmt.Println("Bad Request error", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	data := repository.Bike{
		BikeTypeId:  payload.BikeTypeId,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
	}

	if err := repository.CreateBike(ctx, data); err != nil {
		fmt.Println("Error create bike", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Create Bike"

		ctx.JSON(res.Code, res)
		return
	}
	res.Code = http.StatusOK
	res.Message = "Success Create Bike"

	ctx.JSON(res.Code, res)
}

func GetAllBike(ctx *gin.Context) {
	res := Response{}
	params := models.ReqParams{}

	if err := ctx.ShouldBindQuery(&params); err != nil {
		fmt.Println("Error bind query", err)

		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	data, err := repository.GetAllBike(ctx, params)

	if err != nil {
		fmt.Println("Error get all bike", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Server Error"

		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Success Get Data"
	res.Data = data

	ctx.JSON(res.Code, res)
}
