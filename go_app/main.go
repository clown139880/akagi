package main

import (
	"flag"
	"log"

	c "main/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// The app will run on port 4000 by default, you can custom it with the flag -port
	servePort := flag.String("port", "4000", "Http Server Port")
	flag.Parse()
	log.Printf("server start")

	r := gin.Default()
	r.Use(cors.Default()) //CORS must used before any route
	// Switch to "release" mode in production
	// gin.SetMode(gin.ReleaseMode)
	r.LoadHTMLGlob("views/*")
	// Create a static assets router
	// r.Static("/assets", "./public/assets")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	// Then we bind some route to some handler(controller action)
	r.GET("/", c.HomeHandler)
	r.GET("/posts", c.PostHandler)
	r.GET("/events", c.EventHandler)
	r.POST("/posts", c.CreatePost)
	r.GET("/post/:id", c.ShowPost)
	r.PUT("/post/:id", c.ShowPost)
	v1 := r.Group("/api/v1/weights")
	{
		v1.GET("/", c.FetchAllWeight)
		v1.POST("/", c.CreateWeight)
		v1.GET("/:id", c.FetchSingleWeight)
	}

	// Let's start the server
	r.Run(":" + *servePort)
}
