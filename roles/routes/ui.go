package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{})
}

func LoginForm(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{})
}

func Admin(c *gin.Context) {

	if valid, ok := c.Get("valid"); ok {
		if !valid.(bool) {
			log.Println("session not longer valid pal")
			c.HTML(
				http.StatusOK,
				"login.html",
				gin.H{},
			)
			return
		}
	}

	c.HTML(
		http.StatusOK,
		"admin.html",
		gin.H{})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": "forbidded pal",
	})
}
