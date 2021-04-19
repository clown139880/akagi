package main

import (
	"flag"
	"log"

	c "go_app/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
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
	//r.GET("/posts", c.PostHandler)
	r.POST("/posts", c.CreatePost)
	r.GET("/post/:id", c.ShowPost)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/weights/", c.FetchAllWeight)
		v1.POST("/weights/", c.CreateWeight)
		v1.GET("/weight/:id", c.FetchSingleWeight)

		v1.GET("/photos/", c.FetchAllPhoto)
		v1.GET("/photo/:id", c.FetchSinglePhoto)

		v1.GET("/posts/", c.FetchAllPost)
		v1.POST("/posts/", c.CreatePost)
		v1.GET("/post/:id", c.ShowPost)
		v1.PUT("/post/:id", c.UpdatePost)

		v1.GET("/events/", c.FetchAllEvent)
		v1.GET("/event/:id", c.ShowEvent)
		v1.POST("/events", c.CreateEvent)
		v1.PUT("/event/:id", c.UpdateEvent)
	}

	// Let's start the server
	r.Run(":" + *servePort)
}
