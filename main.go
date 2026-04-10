package main

import (
	"gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		// c.Fail(500, "Internal Server Error")
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v milliseconds for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t).Milliseconds())
	}
}

func main() {
	r := gee.New()
	r.Use(gee.Logger()) // global middleware
	r.GET("/", func(c *gee.Context) {
		time.Sleep(10 * time.Millisecond)
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			time.Sleep(20 * time.Millisecond)
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
