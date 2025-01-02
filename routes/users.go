package routes

import (
	"net/http"

	"example.com/project_1/models"
	"example.com/project_1/utils"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse user data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User Created", "event": user})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse user data"})
		return
	}

	err = user.ValidateUser()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authinticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "token Can't generated"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Login Successful", "token": token})

}
