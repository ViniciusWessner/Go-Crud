package controller

import (
	"fmt"

	"github.com/ViniciusWessner/Go-Crud/src/configuration/rest_err"
	"github.com/ViniciusWessner/Go-Crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil{
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Temos alguns campos incorretos, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}