package routes

import (
	"REST_API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant fetch events",
		})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "cant parse event id ",
		})
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant fetch event",
		})
	}
	context.JSON(http.StatusOK, event)
}

func CreateEvents(context *gin.Context) {

	

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the req",
		})
		return
	}
	userId := context.GetInt64("userId")
	event.UserId = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant save event",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created ", "event": event})
}

// func updateEvent(context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{
// 			"message": "cant parse event id ",
// 		})
// 	}
// 	_, err = models.GetEventById(eventId)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "cant fetch event",
// 		})
// 	}
// 	var updatedEvent models.Event
// 	err = context.ShouldBindJSON(&updatedEvent)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{
// 			"message": "could not parse the req",
// 		})
// 	}
// 	updatedEvent.ID = eventId
// 	err = updatedEvent.Updated()
//  if err != nil{
// 	context.JSON(http.StatusInternalServerError, gin.H{
// 		"message": "cant update event",
// 	})
//  }
// //  context.JSON(http.StatusOK, gin.H{"message":"event updated ok"})

// }
func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "can't parse event id",
		})
		return
	}

	// Fetch existing event by ID
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't fetch event",
		})
		return
	}
	if event.UserId != userId{
		 context.JSON(http.StatusUnauthorized,gin.H{
			"message": "not authorize to update event",
		})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the request",
		})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Updated()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't update event",
		})
		return
	}

	// Return a success response
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "can't parse event id",
		})
		return
	}
	userId := context.GetInt64("userId")

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't fetch event",
		})
		return
	}
	if event.UserId != userId{
		context.JSON(http.StatusUnauthorized,gin.H{
		   "message": "not authorize to delete event",
	   })
	   return
   }
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't fetch event",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "cant delete event"})
}
