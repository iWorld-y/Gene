package main

import (
	"github.com/iWorld-y/EugeneGin/Gene"
	"log"
	"net/http"
	"time"
)

func main() {
	r := Gene.NewEngine()
	r.Use(Gene.Logger)
	r.GET("/", func(c *Gene.Context) {
		c.HTML(http.StatusOK, "<h1>Hello From Gene By iWorld</h1>")
	})
	r.GET("/hello", func(c *Gene.Context) {
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *Gene.Context) {
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.POST("/login", func(c *Gene.Context) {
		c.JSON(http.StatusOK, Gene.H{
			"UserName": c.PostFrom("UserName"),
			"PassWord": c.PostFrom("PassWord"),
		})
	})
	r.GET("/assets/*filepath", func(c *Gene.Context) {
		c.JSON(http.StatusOK, Gene.H{"filepath": c.Param("filepath")})
	})

	v1 := r.Group("/v1")
	v1.GET("/", func(c *Gene.Context) {
		c.HTML(http.StatusOK, "<h1>v1v1\nHello From Gene By iWorld</h1>")
	})
	v1.GET("/hello", func(c *Gene.Context) {
		c.String(http.StatusOK, "v1v1\nHello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	v1.GET("/hello/:name", func(c *Gene.Context) {
		c.String(http.StatusOK, "v1v1\nHello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	v1.POST("/login", func(c *Gene.Context) {
		c.JSON(http.StatusOK, Gene.H{
			"UserName": c.PostFrom("UserName"),
			"PassWord": c.PostFrom("PassWord"),
		})
	})
	v1.GET("/assets/*filepath", func(c *Gene.Context) {
		c.JSON(http.StatusOK, Gene.H{"v1v1\nfilepath": c.Param("filepath")})
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	v2.GET("/hello/:name", func(c *Gene.Context) {
		c.String(http.StatusOK, "v1v1\nHello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	r.Run(":9999")
}

func onlyForV2() Gene.HandlerFunc {
	return func(ctx *Gene.Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
