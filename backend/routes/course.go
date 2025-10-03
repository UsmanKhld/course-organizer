package routes

import (
	"net/http"

	"github.com/UsmanKhld/course-organizer/config"
	"github.com/UsmanKhld/course-organizer/models"
	"github.com/gin-gonic/gin"
)

func CourseRoutes(r *gin.Engine) {
	r.GET("/courses", func(c *gin.Context) {
		rows, err := config.DB.Query(c, "SELECT * FROM courses")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var courses []models.Course
		for rows.Next() {
			var course models.Course
			rows.Scan(&course.ID, &course.UserID, &course.Title, &course.Instructor, &course.Semester)
			courses = append(courses, course)
		}

		c.JSON(http.StatusOK, gin.H{"courses": courses})
	})
}
