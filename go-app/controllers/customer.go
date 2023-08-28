package controllers

import (
	"context"
	"time"

	"github.com/marioheryanto/dbo/dtos"
	"github.com/marioheryanto/dbo/helper"
	"github.com/marioheryanto/dbo/models"
)

type CustomerController struct {
	repo      models.CustomerModelInterface
	validator *helper.Validator
}

type CustomerControllerInterface interface {
	AddCustomer(ctx context.Context, customer dtos.Customer) error
	GetCustomer(ctx context.Context, params dtos.GetCustomersParams) ([]dtos.Customer, error)
	GetCustomerDetail(ctx context.Context, email string) (dtos.Customer, error)
	EditCustomer(ctx context.Context, customer dtos.Customer, email string) error
	DeleteCustomer(ctx context.Context, email string) error
}

func NewCustomerController(repo models.CustomerModelInterface, validator *helper.Validator) CustomerControllerInterface {
	return &CustomerController{
		repo:      repo,
		validator: validator,
	}
}

func (l *CustomerController) AddCustomer(ctx context.Context, customer dtos.Customer) error {
	// validation
	err := l.validator.ValidateStruct(customer)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err = l.repo.CreateCustomer(ctx, customer)

	return err
}

func (l *CustomerController) GetCustomer(ctx context.Context, params dtos.GetCustomersParams) ([]dtos.Customer, error) {
	customers, err := l.repo.GetCustomers(ctx, params)
	if err != nil {
		return customers, err
	}

	return customers, err
}

func (l *CustomerController) GetCustomerDetail(ctx context.Context, email string) (dtos.Customer, error) {
	customer, err := l.repo.GetCustomerDetail(ctx, email)
	if err != nil {
		return customer, err
	}

	return customer, err
}

func (l *CustomerController) EditCustomer(ctx context.Context, customer dtos.Customer, email string) error {
	// validation
	err := l.validator.ValidateStruct(customer)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err = l.repo.EditCustomer(ctx, customer, email)

	return err
}

func (l *CustomerController) DeleteCustomer(ctx context.Context, email string) error {
	_, err := l.repo.DeleteCustomer(ctx, email)
	if err != nil {
		return err
	}

	return err
}
