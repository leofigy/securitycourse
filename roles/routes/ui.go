package routes

import (
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
