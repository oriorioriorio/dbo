package route

import (
	"github.com/gin-gonic/gin"
	"github.com/marioheryanto/dbo/handlers"
	"github.com/marioheryanto/dbo/middlewares"
)

func UserRoutes(app *gin.Engine, c handlers.UserHandlerInterface) {
	// app.Use()
	user := app.Group("/user")
	user.POST("/register", c.Register)
	user.POST("/login", c.Login)
	user.POST("/logout", middlewares.Auth, c.Logout)
	user.GET("/login-data", middlewares.Auth, c.GetLoginData)
}
