package main

import (
	"fmt"
	"roles/model"
	"roles/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	r.GET("/ping", routes.Ping)
	r.GET("/login", routes.LoginForm)
	r.GET("/forbidden", routes.Forbidden)
	r.POST("/forbidden", routes.Forbidden)

	r.POST("/login", routes.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
