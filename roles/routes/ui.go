package routes

import (
	"log"
	"net/http"
	"roles/model"

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
			log.Println("ERROR-> ADMIN PAGE REDIRECTION, SESSION NOT VALID")
			c.HTML(
				http.StatusOK,
				"login.html",
				gin.H{},
			)
			return
		}
	}

	db, err := helperGetDB(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "db-down",
		})
		return
	}

	users := []model.User{}

	result := db.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "db-down",
		})
		return
	}

	c.HTML(
		http.StatusOK,
		"admin.html",
		gin.H{
			"Users": users,
		})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": "forbidded pal",
	})
}
