package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/ViniciusWessner/Go-Crud/src/configuration/rest_err"
)

func NewUserDomain( //construtor do meu objeto userdomain
	email, password, name string,
	age int8,
) UserDomainInterface{
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string 
    Password string 
	Name     string 
	Age      int8   
}

func (ud *UserDomain) EncryptPassword(){ //encripta a senha do usuario
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}