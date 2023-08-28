package controllers

import (
	"context"
	"time"

	"github.com/marioheryanto/dbo/dtos"
	"github.com/marioheryanto/dbo/helper"
	"github.com/marioheryanto/dbo/models"
)

type OrderController struct {
	repo      models.OrderModelInterface
	validator *helper.Validator
}

type OrderControllerInterface interface {
	AddOrder(ctx context.Context, order dtos.Order) error
	GetOrder(ctx context.Context, params dtos.GetOrdersParams) ([]dtos.Order, error)
	GetOrderDetail(ctx context.Context, id interface{}) (dtos.Order, error)
	EditOrder(ctx context.Context, order dtos.Order, id interface{}) error
	DeleteOrder(ctx context.Context, id interface{}) error
}

func NewOrderController(repo models.OrderModelInterface, validator *helper.Validator) OrderControllerInterface {
	return &OrderController{
		repo:      repo,
		validator: validator,
	}
}

func (l *OrderController) AddOrder(ctx context.Context, order dtos.Order) error {
	// validation
	err := l.validator.ValidateStruct(order)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err = l.repo.CreateOrder(ctx, order)

	return err
}

func (l *OrderController) GetOrder(ctx context.Context, params dtos.GetOrdersParams) ([]dtos.Order, error) {
	orders, err := l.repo.GetOrders(ctx, params)
	if err != nil {
		return orders, err
	}

	return orders, err
}

func (l *OrderController) GetOrderDetail(ctx context.Context, id interface{}) (dtos.Order, error) {
	order, err := l.repo.GetOrderDetail(ctx, id)
	if err != nil {
		return order, err
	}

	return order, err
}

func (l *OrderController) EditOrder(ctx context.Context, order dtos.Order, id interface{}) error {
	// validation
	err := l.validator.ValidateStruct(order)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err = l.repo.EditOrder(ctx, order, id)

	return err
}

func (l *OrderController) DeleteOrder(ctx context.Context, id interface{}) error {
	_, err := l.repo.DeleteOrder(ctx, id)
	if err != nil {
		return err
	}

	return err
}
