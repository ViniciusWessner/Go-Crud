package controller

import (
	"net/http"

	"github.com/ViniciusWessner/Go-Crud/src/configuration/logger"
	"github.com/ViniciusWessner/Go-Crud/src/configuration/validation"
	"github.com/ViniciusWessner/Go-Crud/src/controller/model/request"
	"github.com/ViniciusWessner/Go-Crud/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Iniciando CreateUser controller",
		zap.String("Jornada", "CreateUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao validar informacoes do usuario", err,
			zap.String("Jornada", "CreateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusOK, response)
}
