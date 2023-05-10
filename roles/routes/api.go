package routes

import (
	"log"
	"net/http"
	"roles/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	session := sessions.Default(c)
	log.Println("session id", session)

	ok := session.Get("username")

	log.Println("ACCESSING PING ...", ok)

	count := 0

	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()

	log.Println(count)

	c.JSON(http.StatusOK, gin.H{
		"message": session.Get("username"),
		"valid":   v,
	})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)

	log.Println(session.Get("username"))

	db, err := helperGetDB(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "db-down",
		})
	}

	username := c.PostForm("uname")
	password := c.PostForm("psw")

	user := model.User{}
	db.Where("name = ? AND password = ?", username, password).First(&user)

	log.Println(len(user.Name))

	if user.Name == "" {
		log.Println("moving away pal")
		c.Redirect(http.StatusMovedPermanently, "/forbidden")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "in work",
	})

	session.Set("username", user.Name)
	session.Save()

	log.Println("session id", session.ID())
}
