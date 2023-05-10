package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"roles/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
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
	store := memstore.NewStore([]byte("secret"))

	r.Use(
		func(c *gin.Context) {
			uuid := uuid.New()
			c.Set("uuid", uuid.String())
			c.Set("db", persistence)
			c.Next()
		},
		sessions.Sessions("mysession", store),
	)

	r.GET("/ping", func(c *gin.Context) {
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
