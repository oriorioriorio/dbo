package models

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/marioheryanto/dbo/dtos"
	"github.com/marioheryanto/dbo/helper"
)

type UserModel struct {
	dbClient *sql.DB
}

type UserModelInterface interface {
	PingDB() error
	CreateUser(user dtos.User) error
	GetUserWithEmail(user *dtos.User) error
	CheckUserExistWithEmail(email string) (bool, error)
}

func NewUserModel(dbClient *sql.DB) UserModelInterface {
	return &UserModel{
		dbClient: dbClient,
	}
}

func (r *UserModel) PingDB() error {
	err := r.dbClient.Ping()
	if err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	log.Println("dbClient Ping success")

	return nil
}

func (r *UserModel) CreateUser(user dtos.User) error {
	query, args, err := squirrel.Insert("users").Columns("name,email,password").Values(user.Name, user.Email, user.Password).ToSql()
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

func (r *UserModel) GetUserWithEmail(user *dtos.User) error {
	query, args, err := squirrel.Select("name,password").From("users").Where(squirrel.Eq{"email": user.Email}).ToSql()
	if err != nil {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	err = r.dbClient.QueryRow(query, args...).Scan(&user.Name, &user.Password)
	if err != nil && err != sql.ErrNoRows {
		return helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	if err == sql.ErrNoRows {
		return helper.NewServiceError(http.StatusNotFound, "data not found")
	}

	return nil
}

func (r *UserModel) CheckUserExistWithEmail(email string) (bool, error) {
	var userName string

	query, args, err := squirrel.Select("name").From("users").Where(squirrel.Eq{"email": email}).ToSql()
	if err != nil {
		return false, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	err = r.dbClient.QueryRow(query, args...).Scan(&userName)
	if err != nil && err != sql.ErrNoRows {
		return false, helper.NewServiceError(http.StatusInternalServerError, err.Error())
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}
