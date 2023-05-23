package main

import (
	"project/controllers"
	"project/db"
	"project/models"
	"project/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	user := models.User{Email: "foo@gmail.com", FirstName: "Vadym", LastName: "Poliak", Password: "secret", Phone: "1234-44"}
	user.Children = []models.Children{
		models.Children{
			ChildName:      "Susan Doe",
			ChildAge:       4,
			ChildBirthDate: time.Now(),
		},
		models.Children{
			ChildName:      "John Doe",
			ChildAge:       10,
			ChildBirthDate: time.Now(),
		},
	}
	user.Campaigns = []models.Campaign{
		models.Campaign{
			Name:      "Campaign1",
			BeginDate: time.Now(),
			Events: []models.Event{
				{
					Name: "evcp1ev1",
					Date: time.Now(),
					EventItems: []models.EventItem{
						models.EventItem{
							Name:  "cp1ev1evt1",
							Date:  time.Now(),
							Taken: false,
						},
					},
				},
				{
					Name: "evcp1ev2",
					Date: time.Now(),
				},
			},
		},
	}
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
