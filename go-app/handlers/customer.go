package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/dbo/controllers"
	"github.com/marioheryanto/dbo/dtos"
	"github.com/marioheryanto/dbo/helper"
)

type CustomerHandler struct {
	ctrl controllers.CustomerControllerInterface
}

type CustomerHandlerInterface interface {
	AddCustomer(c *gin.Context)
	GetCustomer(c *gin.Context)
	GetCustomerDetail(c *gin.Context)
	EditCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

func NewCustomerHandler(ctrl controllers.CustomerControllerInterface) CustomerHandlerInterface {
	return &CustomerHandler{
		ctrl: ctrl,
	}
}

func (h *CustomerHandler) AddCustomer(c *gin.Context) {
	request := dtos.Customer{}
	response := dtos.Response{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	err = h.ctrl.AddCustomer(context.Background(), request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = "Customer added"
	c.JSON(http.StatusCreated, response)
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	response := dtos.Response{}
	page, _ := strconv.ParseUint(c.Query("page"), 10, 64)
	dataPerPage, _ := strconv.ParseUint(c.Query("data_per_page"), 10, 64)

	customers, err := h.ctrl.GetCustomer(context.Background(), dtos.GetCustomersParams{Page: page, DataPerPage: dataPerPage})
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = customers
	c.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) GetCustomerDetail(c *gin.Context) {
	response := dtos.Response{}

	email := c.Param("email")

	customer, err := h.ctrl.GetCustomerDetail(context.Background(), email)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = customer
	c.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) EditCustomer(c *gin.Context) {
	request := dtos.Customer{}
	response := dtos.Response{}

	email := c.Param("email")

	err := c.Bind(&request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	err = h.ctrl.EditCustomer(context.Background(), request, email)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = "Customer edited"
	c.JSON(http.StatusCreated, response)
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	response := dtos.Response{}

	email := c.Param("email")

	err := h.ctrl.DeleteCustomer(context.Background(), email)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = "Customer deleted"
	c.JSON(http.StatusOK, response)
}
