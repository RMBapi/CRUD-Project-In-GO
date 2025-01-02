package routes

import (
	"net/http"
	"strconv"

	"example.com/project_1/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch events"})
		return
	}
	context.JSON(http.StatusOK, events)

}

func createEvents(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse request data"})
		return
	}

	UId := context.GetInt64("UId")
	event.UserID = UId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})

}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fatch id"})
		return
	}
	event, err := models.GetAllEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch data"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fatch id"})
		return
	}
	UId := context.GetInt64("UId")
	event, err := models.GetAllEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch data"})
		return
	}

	if event.UserID != UId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "The person don't have the permission to edit it"})
		return
	}

	var UpdatedEvent models.Event
	err = context.ShouldBindJSON(&UpdatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse request data"})
		return
	}
	UpdatedEvent.Id = eventID
	err = UpdatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch data"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Updated", "event": UpdatedEvent})

}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fatch id"})
		return
	}
	UId := context.GetInt64("UId")
	event, err := models.GetAllEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fatch data"})
		return
	}

	if event.UserID != UId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "The person don't have the permission to delete it"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted"})

}
