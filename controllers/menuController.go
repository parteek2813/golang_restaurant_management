package controllers

import (
	"context"
	"golang_restaurant_management/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context){

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := menuCollection.Find(context.TODO(), bson.M{})

		defer cancel()

		responseData := gin.H{
			"error":"error occured while listing the menu items",
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, responseData)
		}

		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}


		// if all went well
		c.JSON(http.StatusOK, allMenus)

	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}


func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context){

	}
}






















