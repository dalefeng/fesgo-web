package main

import (
	"fmt"
	"github.com/dalefeng/fesgo"
	"net/http"
)

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
		fmt.Println("te")
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
		}{Name: "Feng "}
		//ctx.HTMLTemplate("login.html", "", "tpl/login.html")
		ctx.HTMLTemplate("index.html", user, "tpl/index.html", "tpl/header.html")
	})
	group.Get("/index", func(ctx *fesgo.Context) {
		user := struct {
			Name string
		}{Name: "Feng "}
		//ctx.HTMLTemplate("login.html", "", "tpl/login.html")
		ctx.HTMLTemplateGlob("index.html", user, "tpl/*.html")
	})
	engine.Run()
}
