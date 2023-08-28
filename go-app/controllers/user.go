package controllers

import (
	"net/http"
	"strings"

	"github.com/marioheryanto/dbo/dtos"
	"github.com/marioheryanto/dbo/helper"
	"github.com/marioheryanto/dbo/models"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Repo      models.UserModelInterface
	LoginRepo models.LoginModelInterface
	Validator *helper.Validator
}

type UserControllerInterface interface {
	Register(user dtos.UserRequest) error
	Login(user dtos.Login) error
	GetLoginData(email string) ([]dtos.LoginData, error)
}

func NewUserController(repo models.UserModelInterface, LoginRepo models.LoginModelInterface, validator *helper.Validator) UserControllerInterface {
	return &UserController{
		Repo:      repo,
		LoginRepo: LoginRepo,
		Validator: validator,
	}
}

func (l *UserController) Register(request dtos.UserRequest) error {
	// validation
	err := l.Validator.ValidateStruct(request)
	if err != nil {
		return err
	}

	user := dtos.User{}
	user.Name = request.Name
	user.Email = strings.ToLower(request.Email)

	exist, err := l.Repo.CheckUserExistWithEmail(request.Email)
	if err != nil {
		return err
	}

	if exist {
		return helper.NewServiceError(http.StatusBadRequest, "email already taken")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		return err
	}

	user.Password = hashPassword

	err = l.Repo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (l *UserController) Login(request dtos.Login) error {
	// validation
	err := l.Validator.ValidateStruct(request)
	if err != nil {
		return err
	}

	user := dtos.User{}
	user.Email = strings.ToLower(request.Email)

	err = l.Repo.GetUserWithEmail(&user)
	if err != nil {
		return helper.ReplaceServiceErrorForLogin(err)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password))
	if err != nil {
		return helper.ReplaceServiceErrorForLogin(err)
	}

	l.LoginRepo.CreateLogin(user.Email)

	return nil
}

func (l *UserController) GetLoginData(email string) ([]dtos.LoginData, error) {

	return l.LoginRepo.GetLoginWithEmail(email)
}
