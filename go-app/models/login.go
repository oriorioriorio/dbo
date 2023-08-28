package models

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/marioheryanto/dbo/dtos"
	"github.com/marioheryanto/dbo/helper"
)

type LoginModel struct {
	dbClient *sql.DB
}

type LoginModelInterface interface {
	PingDB() error
	CreateLogin(email string) error
	GetLoginWithEmail(email string) ([]dtos.LoginData, error)
}

func NewLoginModel(dbClient *sql.DB) LoginModelInterface {
	return &LoginModel{
		dbClient: dbClient,
	}
}

func (r *LoginModel) PingDB() error {
	err := r.dbClient.Ping()
	if err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	log.Println("dbClient Ping success")

	return nil
}

func (r *LoginModel) CreateLogin(email string) error {
	query, args, err := squirrel.Insert("logins").Columns("email").Values(email).ToSql()
	if err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	result, err := r.dbClient.Exec(query, args...)
	if err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	log.Printf("%v", result)

	return nil
}

func (r *LoginModel) GetLoginWithEmail(email string) ([]dtos.LoginData, error) {
	data := []dtos.LoginData{}

	builder := squirrel.Select("last_login").
		From("logins").Where(squirrel.Eq{"email": email})

	query, args, err := builder.ToSql()
	if err != nil {
		return data, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	rows, err := r.dbClient.Query(query, args...)
	if err != nil && err != sql.ErrNoRows {
		return data, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	if err == sql.ErrNoRows {
		return data, helper.NewServiceError(http.StatusNotFound, "data not found")
	}

	defer rows.Close()

	for rows.Next() {
		login := dtos.LoginData{}

		err := rows.Scan(&login.LastLogin)
		if err != nil {
			return data, helper.NewServiceError(http.StatusInternalServerError, err.Error())
		}

		data = append(data, login)
	}

	if err := rows.Err(); err != nil {
		return data, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	return data, nil
}
