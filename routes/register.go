package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sakhdevel/go-web-service/models"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
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
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Could not fetch event"},
		)
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Could not register for an event"},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{"message": "Registered"},
	)
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Could not cancel registration"},
		)
		return
	}
	context.JSON(
		http.StatusOK,
		gin.H{"message": "Cancelled registration"},
	)
}
