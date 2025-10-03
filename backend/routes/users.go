package routes

import (
	"net/http"

	"github.com/UsmanKhld/course-organizer/config"
	"github.com/UsmanKhld/course-organizer/models"
	"github.com/gin-gonic/gin"
)

// GetCourses retrieves all courses for a specific user
func UserRoutes(r *gin.Engine) {
	r.GET("/users", func(c *gin.Context) {
		rows, err := config.DB.Query(c, "SELECT * FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var user models.User
			rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
			users = append(users, user)
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	})
}
