package routes

import (
	"github.com/gin-gonic/gin"

	"example.com/rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
}
