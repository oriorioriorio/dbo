package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/dbo/dtos"
)

func NewServiceError(code int, message string) dtos.ServiceError {
	return dtos.ServiceError{
		Code:    code,
		Message: message,
	}
}

func GenerateResponse(c *gin.Context, r *dtos.Response, err error) (int, dtos.Response) {
	code := http.StatusInternalServerError

	//safety assert
	serviceErr, ok := err.(dtos.ServiceError)
	if !ok {
		r.Message = err.Error()
		return code, *r
	}

	if serviceErr.Code != 0 {
		code = serviceErr.Code
	}

	r.Message = serviceErr.Error()

	return code, *r
}

func ReplaceServiceErrorForLogin(err error) error {
	serviceErr, ok := err.(dtos.ServiceError)
	if serviceErr.Code == http.StatusNotFound || !ok {
		serviceErr.Code = http.StatusBadRequest
		serviceErr.Message = "email atau password salah"
	}

	return serviceErr
}
