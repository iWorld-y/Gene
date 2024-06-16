package main

import (
	"fmt"
	Gene "github.com/iWorld-y/EugeneGin/src"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := Gene.NewEngine()
	r.Use(Gene.Logger)
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("asserts/templates/*")
	//r.Static("/asserts", "static")
	r.Static("/assert", "asserts/static")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *Gene.Context) {
		c.HTML(http.StatusOK, "实时计算技术 环境搭建.html", nil)
	})
	r.GET("/students", func(c *Gene.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", Gene.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *Gene.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", Gene.H{
			"title": "gee",
			"now":   time.Now(),
		})
	})

	r.Run(":9999")
}
