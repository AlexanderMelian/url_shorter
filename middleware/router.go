package middleware

import (
	"localhost/controller"
	"os"

	"github.com/gin-gonic/gin"
)

func Routers() {

	host := os.Getenv("HOST")
	router := gin.Default()

	login := router.Group("/login")
	login.POST("/", Login)

	publicUser := router.Group("/user")
	publicUser.POST("/", controller.CreateUser)

	protectedUser := router.Group("/user")
	protectedUser.Use(Handler())
	protectedUser.GET("/", controller.FindAllUsers)
	//protected.GET("/:id", controller.FindById)

	protectedShorted := router.Group("/shorter")
	protectedShorted.Use(Handler())
	protectedShorted.POST("/", controller.CreateUrlShorted)
	protectedShorted.DELETE("/:url", controller.DeleteUrlShorted)

	redirect := router.Group("/")
	redirect.GET("/:url", controller.RedirectUrl)

	router.Run(host)
}
