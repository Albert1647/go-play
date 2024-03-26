package routes

import (
	"github.com/gin-gonic/gin"
	"natthan.com/go-play/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", getHealth)
	server.GET("/users", getUsers)
	server.GET("/users/:id", getUserByID)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	server.GET("/registration", getRegistrations)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
