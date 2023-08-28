package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marioheryanto/dbo/controllers"
	"github.com/marioheryanto/dbo/database"
	"github.com/marioheryanto/dbo/handlers"
	"github.com/marioheryanto/dbo/helper"
	"github.com/marioheryanto/dbo/models"
	"github.com/marioheryanto/dbo/route"
)

func init() {
	godotenv.Load()
}

func main() {
	// clients
	dbClient := database.ConnectDB()
	validator := helper.NewValidator()

	// models
	userModel := models.NewUserModel(dbClient)
	customerModel := models.NewCustomerModel(dbClient)
	orderModel := models.NewOrderModel(dbClient)
	loginModel := models.NewLoginModel(dbClient)

	// controllers
	userController := controllers.NewUserController(userModel, loginModel, validator)
	customerController := controllers.NewCustomerController(customerModel, validator)
	orderController := controllers.NewOrderController(orderModel, validator)

	// handlers
	userHandler := handlers.NewUserHandler(userController)
	customerHandler := handlers.NewCustomerHandler(customerController)
	orderHandler := handlers.NewOrderHandler(orderController)

	router := gin.Default()

	// config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	// config.AllowHeaders = []string{}

	// ----- Router -----
	router.Use(cors.New(config))

	route.UserRoutes(router, userHandler)
	route.CustomerRoutes(router, customerHandler)
	route.OrderRoutes(router, orderHandler)

	// router.Run(fmt.Sprintf("http://localhost:%v", os.Getenv("PORT")))
	router.Run(":" + os.Getenv("PORT"))

}
