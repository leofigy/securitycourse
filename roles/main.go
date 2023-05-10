package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"roles/model"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func main() {
	persistence := model.NewPersistence("localdb", false, "")
	handler, err := persistence.GetDB()

	/*
		role := model.Role{
			Name:        "admin",
			Description: "super powers",
		}

		user := model.User{
			FullName: "Angel Figueroa",
			Name:     "leofigy",
			Email:    "angel.fig@email.com",
			Password: "welcome1",
			Roles:    []model.Role{role},
		}*/

	//handler.Create(&user)

	if err != nil {
		panic(err)
	}

	fmt.Println(handler)

	r := gin.Default()

	r.LoadHTMLGlob(
		"html/*",
	)

	// session definition
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.Use(
		func(c *gin.Context) {
			log.Println("okay testing again pal")
			uuid := uuid.New()
			c.Set("uuid", uuid.String())
			c.Set("db", persistence)
			fmt.Printf("The request with uuid %s is started \n", uuid)
			c.Next()
			fmt.Printf("The request with uuid %s is served \n", uuid)
		},
	)

	r.GET("/ping", func(c *gin.Context) {
		session := sessions.Default(c)

		ok := session.Get("username")

		log.Println("ACCESSING PING ...", ok)
		/*
			if ok == nil {
				c.Redirect(http.StatusMovedPermanently, "/login")
			}
		*/
		c.JSON(http.StatusOK, gin.H{
			"message": session.Get("username"),
			"valid":   session.Get("last_activity"),
		})

	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"login.html",
			gin.H{})
	})

	r.GET("/forbidden", func(c *gin.Context) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "forbidded pal",
		})
	})

	r.POST("/forbidden", func(c *gin.Context) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "forbidded pal",
		})
	})

	r.POST("/login", func(c *gin.Context) {
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

		session := sessions.Default(c)

		session.Set("username", user.Name)
		session.Set("last_activity", time.Now())
		session.Save()
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func helperGetDB(c *gin.Context) (*gorm.DB, error) {
	obj, ok := c.Get("db")

	if !ok {
		return nil, errors.New("no database available in context")
	}
	provider := obj.(model.Provider)
	return provider.GetDB()
}
