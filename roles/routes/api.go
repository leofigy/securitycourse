package routes

import (
	"log"
	"net/http"
	"roles/model"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

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

	session := sessions.Default(c)
	log.Println("session id", session)
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
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{}
	db.Preload("Roles").Where("name = ? AND password = ?", username, password).First(&user)

	if user.Name == "" {
		log.Println("redirecting to forbidden")
		c.Redirect(http.StatusMovedPermanently, "/forbidden")
		return
	}

	session.Set("username", user.Name)
	session.Set("active", time.Now().Unix())
	session.Save()
	log.Println("session id", session.ID())

	if len(user.Roles) > 0 {
		// just grab the first role
		switch user.Roles[0].Name {
		case "admin":
			c.Redirect(
				http.StatusMovedPermanently, "/admin",
			)
			return
		default:
			c.JSON(http.StatusOK, gin.H{
				"message": "in but not available function",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome!",
		})
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	log.Println("TRYING TO GO OUTSIDE .....")
	log.Println("current user ->", session.Get("username"))
	log.Println("session id", session.ID())
	session.Set("username", "") // this will mark the session as "written" and hopefully remove the username
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "bye!!!",
	})
}
