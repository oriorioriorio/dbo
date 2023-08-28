package route

import (
	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/dbo/handlers"
	"github.com/marioheryanto/dbo/middlewares"
)

func OrderRoutes(app *gin.Engine, c handlers.OrderHandlerInterface) {
	// app.Use()
	order := app.Group("/order")
	order.POST("/", middlewares.Auth, c.AddOrder)
	order.GET("/", middlewares.Auth, c.GetOrder)
	order.GET("/:id", middlewares.Auth, c.GetOrderDetail)
	order.PUT("/:id", middlewares.Auth, c.EditOrder)
	order.DELETE("/:id", middlewares.Auth, c.DeleteOrder)
}
