package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sakhdevel/go-web-service/db"
	"sakhdevel/go-web-service/models"
	"strconv"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8085") // localhost

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch events. Try again later"},
		)
		return
	}
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse event Id"},
		)
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch event"},
		)
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse request data"},
		)
		return
	}

	// todo change later
	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not create event. Try again later"},
		)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
