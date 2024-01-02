package main

import (
	"gym/controllers"

	"github.com/gin-gonic/gin"

	"gym/db"
)

func main() {
	db.Connect()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.Index)
	r.POST("/exercise", controllers.CreateExercise)
	r.POST("/set", controllers.CreateSet)
	r.Run()
}
