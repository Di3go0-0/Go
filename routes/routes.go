package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func MainRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hi! " + name,
		})
	})

	r.POST("/register", func(context *gin.Context) {
		var newUser User

		if err := context.BindJSON(&newUser); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Error to bind JSON",
			})
			return
		}
		if newUser.Name == "" || newUser.Email == "" {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Name and Email are required",
			})
			return
		}

		users = append(users, newUser)

		context.JSON(http.StatusCreated, gin.H{
			"message": "User created",
			"data":    users,
		})
	})

	return r
}