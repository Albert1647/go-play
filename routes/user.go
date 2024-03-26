package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"natthan.com/go-play/models"
	"natthan.com/go-play/utils"
)

func getUsers(context *gin.Context) {
	users, err := models.GetAllUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data"})
		return
	}
	context.JSON(http.StatusOK, users)
}

func getUserByID(context *gin.Context) {
	userID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Parameter"})
		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user or user not found"})
		return
	}
	context.JSON(http.StatusOK, user)
}

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credential"})
		return
	}

	accessToken, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not sign jwt"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "accessToken": accessToken})
}
