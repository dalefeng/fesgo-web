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
	engine := fesgo.NewEngine("8111")
	group := engine.Group("user")

	group.Use(func(next fesgo.HandlerFunc) fesgo.HandlerFunc {
		return func(ctx *fesgo.Context) {
			fmt.Println("pre handler")
			next(ctx)
			fmt.Println("post handler")
		}
	})
	group.Get("/info", func(ctx *fesgo.Context) {
		ctx.W.Write([]byte("get info"))
	}, func(next fesgo.HandlerFunc) fesgo.HandlerFunc {
		return func(ctx *fesgo.Context) {
			fmt.Println("info")
			next(ctx)
			fmt.Println("info post")
		}
	})
	group.Post("/info", func(ctx *fesgo.Context) {
		fmt.Println("test ")
		ctx.W.Write([]byte("pots info"))
	})
	group.Post("/login", func(ctx *fesgo.Context) {
		ctx.W.Write([]byte("login"))
	})
	group.Any("/any", func(ctx *fesgo.Context) {
		ctx.W.Write([]byte("any"))
	})
	group.Get("/get/:id", func(ctx *fesgo.Context) {
		ctx.W.Write([]byte("/get/:id"))
	})
	group.Get("/html", func(ctx *fesgo.Context) {
		ctx.HTML(http.StatusOK, "<h1>FesGO<h1>")
	})
	group.Get("/indexTemplate", func(ctx *fesgo.Context) {
		user := struct {
			Name string
		}{Name: "Feng 4442"}
		//ctx.HTMLTemplate("login.html", "", "tpl/login.html")
		ctx.HTMLTemplate("index.html", user, "tpl/index.html", "tpl/header.html")
	})
	engine.LoadTemplate("tpl/*.html")
	group.Get("/index", func(ctx *fesgo.Context) {
		user := struct {
			Name string
		}{Name: "Feng 123316qq44465"}
		//ctx.HTMLTemplate("login.html", "", "tpl/login.html")
		ctx.Template("index.html", user)
	})
	group.Get("/xml", func(ctx *fesgo.Context) {
		u := &User{Name: "feng"}
		ctx.XML(http.StatusOK, u)
	})
	group.Get("/file", func(ctx *fesgo.Context) {
		ctx.FileAttachment("tmp/main.exe", "man.exe")
	})
	group.Get("/string", func(ctx *fesgo.Context) {
		ctx.String(http.StatusOK, "string")
	})
	group.Get("/redirect", func(ctx *fesgo.Context) {
		ctx.Render(http.StatusFound, &render.Redirect{
			Code:     http.StatusFound,
			Request:  ctx.R,
			Location: "/user/index",
		})
	})
	fmt.Println("server run ...")
	engine.Run()
}
