package middleware

import (
	"localhost/controller"

	"github.com/gin-gonic/gin"
)

func Routers() {
	router := gin.Default()

	login := router.Group("/login")
	login.POST("/", Login)

	publicUser := router.Group("/user")
	publicUser.POST("/", controller.CreateUser)

	protectedUser := router.Group("/user")
	protectedUser.Use(Handler())
	protectedUser.GET("/", controller.FindAllUsers)
	//protected.GET("/:id", controller.FindById)

	protectedShorted := router.Group("/")
	protectedShorted.Use(Handler())
	protectedShorted.POST("/", controller.CreateUrlShorted)

	router.Run("localhost:8080")
}
