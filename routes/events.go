package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sakhdevel/go-web-service/models"
	"sakhdevel/go-web-service/utils"
	"strconv"
)

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
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{"message": "Not authorized"},
		)
		return
	}

	err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{"message": "Not authorized"},
		)
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse request data"},
		)
		return
	}

	// todo change later
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

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse event Id"},
		)
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch the event"},
		)
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse request data"},
		)
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not update event"},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
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
			gin.H{"message": "Could not fetch the event"},
		)
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not delete the event"},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}
