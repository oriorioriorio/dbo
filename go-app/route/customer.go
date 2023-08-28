package route

import (
	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/dbo/handlers"
	"github.com/marioheryanto/dbo/middlewares"
)

func CustomerRoutes(app *gin.Engine, c handlers.CustomerHandlerInterface) {
	// app.Use()
	customer := app.Group("/customer")
	customer.POST("/", middlewares.Auth, c.AddCustomer)
	customer.GET("/", middlewares.Auth, c.GetCustomer)
	customer.GET("/:email", middlewares.Auth, c.GetCustomerDetail)
	customer.PUT("/:email", middlewares.Auth, c.EditCustomer)
	customer.DELETE("/:email", middlewares.Auth, c.DeleteCustomer)
}
