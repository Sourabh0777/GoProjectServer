package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"example.com/rest-api/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the events list."})
	}
	context.JSON(http.StatusOK, events)
}
func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
func getEvent(context *gin.Context) {
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event."})
		return
	}
	event, err := models.GetEventByID(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event."})
	}
	context.JSON(http.StatusOK, event)
}
func updateEvent(context *gin.Context) {
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event."})
		return
	}
	_, err = models.GetEventByID(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event."})
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}
	updatedEvent.ID = eventid
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully."})
}

func deleteEvent(context *gin.Context) {
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event."})
		return
	}
	event, err := models.GetEventByID(eventid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully."})

}
