package handlers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/dbo/controllers"
	"github.com/marioheryanto/dbo/dtos"
	"github.com/marioheryanto/dbo/helper"
)

type UserHandler struct {
	Controller controllers.UserControllerInterface
}

type UserHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	GetLoginData(c *gin.Context)
}

func NewUserHandler(controllers controllers.UserControllerInterface) UserHandlerInterface {
	return &UserHandler{
		Controller: controllers,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	request := dtos.UserRequest{}
	response := dtos.Response{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	err = h.Controller.Register(dtos.UserRequest(request))
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = "Register success"
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) Login(c *gin.Context) {
	request := dtos.Login{}
	response := dtos.Response{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	err = h.Controller.Login(dtos.Login(request))
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	ttl := time.Now().Add(7 * 24 * time.Hour)
	jwt, err := helper.GenerateJWT(jwt.MapClaims{
		"sub": request.Email,
		"exp": ttl.Unix(),
	})
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	helper.SetCookie(c, "jwt", jwt, ttl.Second(), "", "", false, true)

	response.Data = "logged in"
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Logout(c *gin.Context) {
	response := dtos.Response{}
	tokenString, _ := c.Cookie("jwt")

	if tokenString == "" {
		response.Message = "please login first"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	helper.SetCookie(c, "jwt", "", -1, "", "", false, true)

	response.Data = "logged out"
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetLoginData(c *gin.Context) {
	response := dtos.Response{}
	email, _ := c.Get("email")

	data, err := h.Controller.GetLoginData(email.(string))
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = data
	c.JSON(http.StatusOK, response)
}
