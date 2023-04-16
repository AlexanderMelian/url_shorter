package middleware

import (
	"localhost/controller"

	"github.com/gin-gonic/gin"
)

func Routers() {
	router := gin.Default()

	login := router.Group("/login")
	login.POST("/", Login)

	public := router.Group("/user")
	public.GET("/", controller.FindAll)
	public.POST("/", controller.Create)

	protected := router.Group("/user")
	protected.Use(Handler())
	protected.GET("/:id", controller.FindById)

	router.Run("localhost:8080")
}
