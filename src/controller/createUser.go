package controller

import (
	"net/http"

	"github.com/ViniciusWessner/Go-Crud/src/configuration/logger"
	"github.com/ViniciusWessner/Go-Crud/src/configuration/validation"
	"github.com/ViniciusWessner/Go-Crud/src/controller/model/request"
	"github.com/ViniciusWessner/Go-Crud/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(
		userRequest.Email, 
		userRequest.Password, 
		userRequest.Name, 
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Usuario criado com sucesso!",
		zap.String("Jornada", "CreateUser"))

	c.String(http.StatusOK,"")
}
