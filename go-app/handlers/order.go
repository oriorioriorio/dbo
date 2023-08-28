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

type OrderHandler struct {
	ctrl controllers.OrderControllerInterface
}

type OrderHandlerInterface interface {
	AddOrder(c *gin.Context)
	GetOrder(c *gin.Context)
	GetOrderDetail(c *gin.Context)
	EditOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
}

func NewOrderHandler(ctrl controllers.OrderControllerInterface) OrderHandlerInterface {
	return &OrderHandler{
		ctrl: ctrl,
	}
}

func (h *OrderHandler) AddOrder(c *gin.Context) {
	request := dtos.Order{}
	response := dtos.Response{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	err = h.ctrl.AddOrder(context.Background(), request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = "Order added"
	c.JSON(http.StatusCreated, response)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	response := dtos.Response{}
	page, _ := strconv.ParseUint(c.Query("page"), 10, 64)
	dataPerPage, _ := strconv.ParseUint(c.Query("data_per_page"), 10, 64)

	customers, err := h.ctrl.GetOrder(context.Background(), dtos.GetOrdersParams{Page: page, DataPerPage: dataPerPage})
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = customers
	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) GetOrderDetail(c *gin.Context) {
	response := dtos.Response{}
	id := c.Param("id")

	// if id == "" {
	// 	c.JSON(helper.GenerateResponse(c, &response, helper.NewServiceError(http.StatusBadRequest, "order id is empty")))
	// 	return
	// }

	customer, err := h.ctrl.GetOrderDetail(context.Background(), id)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = customer
	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) EditOrder(c *gin.Context) {
	request := dtos.Order{}
	response := dtos.Response{}

	id := c.Param("id")

	err := c.Bind(&request)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	err = h.ctrl.EditOrder(context.Background(), request, id)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = "Order edited"
	c.JSON(http.StatusCreated, response)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	response := dtos.Response{}

	id := c.Param("id")

	err := h.ctrl.DeleteOrder(context.Background(), id)
	if err != nil {
		c.JSON(helper.GenerateResponse(c, &response, err))
		return
	}

	response.Data = "Order deleted"
	c.JSON(http.StatusOK, response)
}
