package controller

import (
	"localhost/models"
	"localhost/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserController(c *gin.Context) {
	var user models.User

	c.JSON(200, user)
}

func FindAll(c *gin.Context) {

	users, flag := service.FindAll()

	if !flag {
		c.JSON(http.StatusOK, gin.H{"users": "Not found"})
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func FindById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "BAD REQUEST")
	}

	user, flag := service.FindById(uint(id))

	if !flag {
		c.String(http.StatusOK, "NOT FOUND")
	}

	c.JSON(http.StatusOK, gin.H{"users": user})
}

func Create(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Json")
		return
	}
	if input.Username == "" || input.Password == "" || input.Name != "" {
		c.JSON(http.StatusBadRequest, "Input error")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Input error")
		return
	}

	//user := models.User{Username: input.Username, Name: input.Name, LastName: input.LastName, Password: string(hashedPassword)}
	input.Password = string(hashedPassword)
	service.CreateUser(input)

	c.JSON(http.StatusOK, "")
}
