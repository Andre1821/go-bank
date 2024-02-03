package main

import (
	"bank.com/controllers"
	"bank.com/initializers"
	"bank.com/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariabels()
	initializers.ConnectToDB()
	initializers.SysncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup )
	r.POST("/login", controllers.Login )
	r.GET("/validate", middleware.RequiredAuth, controllers.Validate )
	r.Run() 
}