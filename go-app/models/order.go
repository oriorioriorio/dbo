package models

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/marioheryanto/dbo/dtos"
	"github.com/marioheryanto/dbo/helper"
)

type OrderModel struct {
	dbClient *sql.DB
}

type OrderModelInterface interface {
	PingDB() error
	CreateOrder(ctx context.Context, params dtos.Order) (interface{}, error)
	GetOrders(ctx context.Context, params dtos.GetOrdersParams) ([]dtos.Order, error)
	GetOrderDetail(ctx context.Context, id interface{}) (dtos.Order, error)
	EditOrder(ctx context.Context, params dtos.Order, id interface{}) (interface{}, error)
	DeleteOrder(ctx context.Context, id interface{}) (int64, error)
}

func NewOrderModel(dbClient *sql.DB) OrderModelInterface {
	return &OrderModel{
		dbClient: dbClient,
	}
}

func (r *OrderModel) PingDB() error {
	err := r.dbClient.Ping()
	if err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *OrderModel) CreateOrder(ctx context.Context, params dtos.Order) (interface{}, error) {
	query, args, err := squirrel.Insert("orders").Columns("name, user_email").
		Values(params.Name, params.UserEmail).
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

func (r *OrderModel) GetOrders(ctx context.Context, params dtos.GetOrdersParams) ([]dtos.Order, error) {
	orders := []dtos.Order{}

	builder := squirrel.Select("name, user_email").
		From("orders").Offset(params.Page)

	builder = builder.Limit(10)
	if params.DataPerPage > 0 {
		builder = builder.Limit(params.DataPerPage)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return orders, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	rows, err := r.dbClient.QueryContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return orders, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	if err == sql.ErrNoRows {
		return orders, helper.NewServiceError(http.StatusNotFound, "data not found")
	}

	defer rows.Close()

	for rows.Next() {
		order := dtos.Order{}

		err := rows.Scan(&order.Name, &order.UserEmail)
		if err != nil {
			return orders, helper.NewServiceError(http.StatusInternalServerError, err.Error())
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return orders, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	return orders, nil
}

func (r *OrderModel) GetOrderDetail(ctx context.Context, id interface{}) (dtos.Order, error) {
	orders := dtos.Order{}

	query, args, err := squirrel.Select("name, user_email").
		From("orders").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return orders, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	err = r.dbClient.QueryRow(query, args...).Scan(&orders.Name, &orders.UserEmail)
	if err != nil && err != sql.ErrNoRows {
		return orders, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	if err == sql.ErrNoRows {
		return orders, helper.NewServiceError(http.StatusNotFound, "data not found")
	}

	return orders, nil
}

func (r *OrderModel) EditOrder(ctx context.Context, params dtos.Order, id interface{}) (interface{}, error) {
	query, args, err := squirrel.Update("orders").
		Set("name", params.Name).
		Set("user_email", params.UserEmail).
		Where(squirrel.Eq{"id": id}).
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

func (r *OrderModel) DeleteOrder(ctx context.Context, id interface{}) (int64, error) {
	query, args, err := squirrel.Delete("orders").Where(squirrel.Eq{"id": id}).ToSql()
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
