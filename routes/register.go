package routes

import (
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func createRegistration(context *gin.Context) {
	userId := context.GetInt64("user_id")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func deleteRegistration(context *gin.Context) {
	userId := context.GetInt64("user_id")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.DeleteRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration deleted"})

}
