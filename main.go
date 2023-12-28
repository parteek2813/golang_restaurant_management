package main

import (
	"golang_restaurant_management/database"
	"golang_restaurant_management/middleware"
	"golang_restaurant_management/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main(){
	port := os.Getenv("PORT")


	// if port nil, then select port 8000
	if port == ""{
		port = "8000"
	}


	// gin router
	router := gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)


	router.Run(":" + port)

}



