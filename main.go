package main

import (
	"github.com/iWorld-y/EugeneGin/Gene"
	"net/http"
)

func main() {
	r := Gene.NewEngine()
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
	r.Run(":9999")
}
