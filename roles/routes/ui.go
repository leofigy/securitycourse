package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginForm(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": "forbidded pal",
	})
}
