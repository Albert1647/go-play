package routes

import (
	"github.com/gin-gonic/gin"
	"natthan.com/go-play/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	v1 := server.Group("/v1")
	{
		v1.GET("/health", getHealth)
		v1.GET("/users", getUsers)
		v1.GET("/users/:id", getUserByID)

		v1.GET("/events", getEvents)
		v1.GET("/events/:id", getEventByID)

		v1.GET("/registration", getRegistrations)

		v1.POST("/signup", signup)
		v1.POST("/login", login)
	}
	authenticated := v1.Group("").Use(middlewares.Authenticate)
	{
		authenticated.POST("/events", createEvent)
		authenticated.PUT("/events/:id", updateEvent)
		authenticated.DELETE("/events/:id", deleteEvent)
		authenticated.POST("/events/:id/register", registerForEvent)
		authenticated.DELETE("/events/:id/register", cancelRegistration)
	}
	v2 := server.Group("/v2")
	{
		v2.GET("/health", getHealth)
		v2.GET("/users", getUsers)
		v2.GET("/users/:id", getUserByID)

		v2.GET("/events", getEvents)
		v2.GET("/events/:id", getEventByID)

		v2.GET("/registration", getRegistrations)

		v2.POST("/signup", signup)
		v2.POST("/login", login)
	}
}
