package routes

import (
	"REST_API/models"
	"REST_API/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse data",
		})
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{"message": "cant save user"})
	}
	context.JSON(http.StatusCreated,gin.H{"message": "user created success"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot parse data",
		})
		return
	}

	// Validate credentials
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Cannot authenticate user",
		})
		return
	}

      token , err := utils.GenerateToken(user.Email,user.ID)
       if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot authenticate user",
		})
		return
	   }

	
	context.JSON(http.StatusOK, gin.H{
		"message": "Login success", "token:": token ,
		
	})
}
