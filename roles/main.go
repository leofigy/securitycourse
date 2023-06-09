package main

import (
	"fmt"
	"log"
	"roles/model"
	"roles/routes"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	persistence := model.NewPersistence("localdb", false, "")
	handler, err := persistence.GetDB()

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
		func(c *gin.Context) {
			log.Println("***** VALIDATION ******")
			session := sessions.Default(c)
			log.Println("*** CURRENT SESSION", session)
			ok := session.Get("username")
			log.Println("ACTIVE USER ->", ok)
			if ok == nil {
				c.Set("valid", false)
			} else {
				log.Println("inspecting time")
				lastUsage := session.Get("active")
				if lastUsage == nil {
					// invalidating session no last_usage info
					c.Set("valid", false)
				}

				utime := lastUsage.(int64)

				workingTime := time.Unix(utime, 0)
				log.Println("ACTIVE TIME ->", workingTime)

				current := time.Now()
				if current.After(workingTime.Add(time.Second * 300)) {
					log.Println("WORKING TIME -> exceeded")
					c.Set("valid", false)
				}
			}
			c.Next()
		},
	)

	r.GET("/", routes.Index)
	r.GET("/login", routes.LoginForm)
	r.GET("/forbidden", routes.Forbidden)
	r.POST("/forbidden", routes.Forbidden)
	r.POST("/login", routes.Login)
	r.GET("/logout", routes.Logout)
	r.POST("/admin", routes.Admin)
	r.GET("/admin", routes.Admin)
	r.GET("/ping", routes.Ping)
	r.POST("/add_user", routes.AddUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
