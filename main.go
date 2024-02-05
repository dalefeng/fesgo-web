package main

import (
	"fmt"
	"github.com/dalefeng/fesgo"
	"github.com/dalefeng/fesgo/render"
	"net/http"
)

type User struct {
	Name string `xml:"name"`
}

func main() {
	engine := fesgo.NewEngine()
	group := engine.Group("user")

	group.Use(func(next fesgo.HandlerFunc) fesgo.HandlerFunc {
		return func(c *fesgo.Context) {
			fmt.Println("pre handler1")
			next(c)
			fmt.Println("post handler")
		}
	})
	group.Get("/info", func(c *fesgo.Context) {
		c.W.Write([]byte("get info"))
	}, func(next fesgo.HandlerFunc) fesgo.HandlerFunc {
		return func(c *fesgo.Context) {
			fmt.Println("info")
			next(c)
			fmt.Println("info post")
		}
	})
	group.Post("/info", func(c *fesgo.Context) {
		fmt.Println("test ")
		c.W.Write([]byte("pots info"))
	})
	group.Post("/login", func(c *fesgo.Context) {
		c.W.Write([]byte("login"))
	})
	group.Any("/any", func(c *fesgo.Context) {
		c.W.Write([]byte("any"))
	})
	group.Get("/get/:id", func(c *fesgo.Context) {
		c.W.Write([]byte("/get/:id"))
	})
	group.Get("/html", func(c *fesgo.Context) {
		c.HTML(http.StatusOK, "<h1>FesGO<h1>")
	})
	group.Get("/indexTemplate", func(c *fesgo.Context) {
		user := struct {
			Name string
		}{Name: "Feng 4442"}
		//c.HTMLTemplate("login.html", "", "tpl/login.html")
		c.HTMLTemplate("index.html", user, "tpl/index.html", "tpl/header.html")
	})
	engine.LoadTemplate("tpl/*.html")
	group.Get("/index", func(c *fesgo.Context) {
		user := struct {
			Name string
		}{Name: "Feng 123316qq44465"}
		//c.HTMLTemplate("login.html", "", "tpl/login.html")
		c.Template("index.html", user)
	})
	group.Get("/xml", func(c *fesgo.Context) {
		u := &User{Name: "feng"}
		c.XML(http.StatusOK, u)
	})
	group.Get("/redirect", func(c *fesgo.Context) {
		c.Render(http.StatusFound, &render.Redirect{
			Code:     http.StatusFound,
			Request:  c.R,
			Location: "/user/index",
		})
	})
	group.Get("/file", func(c *fesgo.Context) {
		c.FileAttachment("tmp/main.exe", "man.exe")
	})
	group.Get("/string", func(c *fesgo.Context) {
		name := c.GetQuery("name")
		ids, _ := c.GetQueryMap("user")
		c.String(http.StatusOK, "string hello: %s, ids: %v", name, ids)
	})

	fmt.Println("server run ...")
	engine.Run()
}
