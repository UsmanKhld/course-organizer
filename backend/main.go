package main

import (
	"github.com/UsmanKhld/course-organizer/config"
	"github.com/UsmanKhld/course-organizer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB() // ✅ initialize the Postgres connection

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.CourseRoutes(r)
	routes.UserRoutes(r)
	r.Run(":8080")
}
