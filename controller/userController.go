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

func FindAllUsers(c *gin.Context) {

	users, err := service.FindAllUsers()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"users": "Not found"})
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func FindUserById(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "BAD REQUEST")
	}

	user, err := service.FindUserById(uint(id))

	if err != nil {
		c.String(http.StatusOK, "NOT FOUND")
	}

	c.JSON(http.StatusOK, gin.H{"users": user})
}

func CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Json")
		return
	}
	if input.Username == "" || input.Password == "" || input.Name == "" || input.Email == "" {
		c.JSON(http.StatusBadRequest, "Input error")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		println("HASH ERROR")
		c.JSON(http.StatusBadRequest, "Input error")
		return
	}

	//user := models.User{Username: input.Username, Name: input.Name, LastName: input.LastName, Password: string(hashedPassword)}
	input.Password = string(hashedPassword)
	err = service.CreateUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Input error")
		return
	}

	c.JSON(http.StatusOK, "User created")
}
