package model

import (
	"fmt"

	"github.com/ViniciusWessner/Go-Crud/src/configuration/logger"
	"github.com/ViniciusWessner/Go-Crud/src/configuration/rest_err"
	"go.uber.org/zap"
)


func (ud *UserDomain) CreateUser() *rest_err.RestErr{

	logger.Info("Iniciando create user MODEL", zap.String("Jornada", "CreateUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}