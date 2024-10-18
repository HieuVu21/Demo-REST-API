package routes

import (
	"REST_API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context){
userId := context.GetInt64("userId")
eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "can't parse event id",
		})
		return
	}

	event,err := models.GetEventById(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"cant fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"cant regis user for event"})
	}
	context.JSON(http.StatusCreated,gin.H{"message":"registered"})
}

func cancelRegistration(context *gin.Context){
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	var event models.Event
	event.ID = eventId
	event.CancelRegistration(userId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"cant cancel regis user for event"})
	}
	context.JSON(http.StatusCreated,gin.H{"message":"cancel registered"})
	
}