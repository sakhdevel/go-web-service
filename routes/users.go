package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sakhdevel/go-web-service/models"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse request data"},
		)
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not save user"},
		)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
