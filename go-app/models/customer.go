package models

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/marioheryanto/dbo/dtos"
	"github.com/marioheryanto/dbo/helper"
)

type CustomerModel struct {
	dbClient *sql.DB
}

type CustomerModelInterface interface {
	PingDB() error
	CreateCustomer(ctx context.Context, params dtos.Customer) (interface{}, error)
	GetCustomers(ctx context.Context, params dtos.GetCustomersParams) ([]dtos.Customer, error)
	GetCustomerDetail(ctx context.Context, email string) (dtos.Customer, error)
	EditCustomer(ctx context.Context, params dtos.Customer, email string) (interface{}, error)
	DeleteCustomer(ctx context.Context, email string) (int64, error)
}

func NewCustomerModel(dbClient *sql.DB) CustomerModelInterface {
	return &CustomerModel{
		dbClient: dbClient,
	}
}

func (r *CustomerModel) PingDB() error {
	err := r.dbClient.Ping()
	if err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *CustomerModel) CreateCustomer(ctx context.Context, params dtos.Customer) (interface{}, error) {
	query, args, err := squirrel.Insert("customers").Columns("name, email, phone").
		Values(params.Name, params.Email, params.Phone).
		ToSql()

	if err != nil {
		return 0, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	result, err := r.dbClient.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	return result.LastInsertId()
}

func (r *CustomerModel) GetCustomers(ctx context.Context, params dtos.GetCustomersParams) ([]dtos.Customer, error) {
	customers := []dtos.Customer{}

	var limit uint64 = 10
	if params.DataPerPage > 0 {
		limit = params.DataPerPage
	}

	if params.Page <= 0 {
		params.Page = 1
	}

	offset := (params.Page * limit) - limit

	builder := squirrel.Select("name, email, phone").
		From("customers").Offset(offset).Limit(limit)

	query, args, err := builder.ToSql()
	if err != nil {
		return customers, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	rows, err := r.dbClient.QueryContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return customers, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	if err == sql.ErrNoRows {
		return customers, helper.NewServiceError(http.StatusNotFound, "data not found")
	}

	defer rows.Close()

	for rows.Next() {
		customer := dtos.Customer{}

		err := rows.Scan(&customer.Name, &customer.Email, &customer.Phone)
		if err != nil {
			return customers, helper.NewServiceError(http.StatusInternalServerError, err.Error())
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return customers, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	return customers, nil
}

func (r *CustomerModel) GetCustomerDetail(ctx context.Context, email string) (dtos.Customer, error) {
	customers := dtos.Customer{}

	query, args, err := squirrel.Select("name, email, phone").
		From("customers").
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		return customers, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	err = r.dbClient.QueryRow(query, args...).Scan(&customers.Name, &customers.Email, &customers.Phone)
	if err != nil && err != sql.ErrNoRows {
		return customers, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	if err == sql.ErrNoRows {
		return customers, helper.NewServiceError(http.StatusNotFound, "data not found")
	}

	return customers, nil
}

func (r *CustomerModel) EditCustomer(ctx context.Context, params dtos.Customer, email string) (interface{}, error) {
	query, args, err := squirrel.Update("customers").
		Set("name", params.Name).
		Set("email", params.Email).
		Set("phone", params.Phone).
		Where(squirrel.Eq{"email": email}).
		ToSql()

	if err != nil {
		return 0, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	result, err := r.dbClient.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	row, _ := result.RowsAffected()
	if row == 0 {
		return 0, helper.NewServiceError(http.StatusBadRequest, "data not found")
	}

	return row, nil
}

func (r *CustomerModel) DeleteCustomer(ctx context.Context, email string) (int64, error) {
	query, args, err := squirrel.Delete("customers").Where(squirrel.Eq{"email": email}).ToSql()
	if err != nil {
		return 0, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	result, err := r.dbClient.ExecContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return 0, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	row, _ := result.RowsAffected()
	if row == 0 {
		return 0, helper.NewServiceError(http.StatusBadRequest, "data not found")
	}

	return row, err
}
