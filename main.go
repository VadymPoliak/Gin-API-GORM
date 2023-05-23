package main

import (
	"project/controllers"
	"project/db"
	"project/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	//var client *mongo.Client

	//var JwtHandlers = jwthandlers.NewJwtHndlrs(client)

	routergin := gin.Default()
	routergin.OPTIONS("/", func(c *gin.Context) {})
	v1 := routergin.Group("/v1")
	tasks := v1.Group("/tasks")

	err := routergin.SetTrustedProxies(nil)
	if err != nil {
		utils.Logger.Printf("error setting proxies : %s", err.Error())
	}

	//router.GET("/users", GetUsers)
	db.ConnectDatabase()
	//tasks.Use(JwtHandlers.AuthRequired)
	tasks.GET("/users", func(ctx *gin.Context) {
		controllers.GetUsers(ctx)
	})
	tasks.POST("/users", func(ctx *gin.Context) {
		controllers.CreateUser(ctx)
	})

	err = routergin.Run("127.0.0.1:8000")
	if err != nil {
		return
	}

}
