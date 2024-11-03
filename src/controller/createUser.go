package controller

import (
	"github.com/ViniciusWessner/Go-Crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	err := rest_err.NewBadRequestError("Voce chamou a rota errada")
	c.JSON(err.Code, err)
}