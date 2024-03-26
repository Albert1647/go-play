package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"natthan.com/go-play/models"
)

func getRegistrations(context *gin.Context) {
	regs, err := models.GetAllRegistration()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get registration", "err": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, regs)
}

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Parameter"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for the event", "err": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Register successfully"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Parameter"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for the event", "err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Delete Registration successfully"})
}
