package routes

import (
	"net/http"
	"strconv"

	"example.com/project_1/models"
	"github.com/gin-gonic/gin"
)

func registerForEvents(context *gin.Context) {
	UId := context.GetInt64("UId")
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

	err = event.Register(UId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Register for User"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Successfully register!!"})

}

func cancelRegistration(context *gin.Context) {
	UId := context.GetInt64("UId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fatch id"})
		return
	}

	var event models.Event
	event.Id = eventID

	err = event.DeleteRegistration(UId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel the request"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Cancel registration successfully!!"})

}
