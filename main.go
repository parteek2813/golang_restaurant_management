package main

import (
	"os"
)

func main(){
	port := os.Getenv("PORT")


	// if port nil, then select port 8000
	if port == ""{
		port = "8000"
	}


	// gin router
	router := gin.New()
	router.Use(gin.Logger())
	router
}



