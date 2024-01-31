package main

import (
	"fmt"
	"github.com/dalefeng/fesgo"
)

func main() {
	engine := fesgo.NewEngine("8111")
	group := engine.Group("user")

	group.PreHandle(func(next fesgo.HandlerFunc) fesgo.HandlerFunc {
		return func(ctx *fesgo.Context) {
			fmt.Println("pre handler")
			next(ctx)
		}
	})
	group.PostHandle(func(next fesgo.HandlerFunc) fesgo.HandlerFunc {
		return func(ctx *fesgo.Context) {
			fmt.Println("post handler")
		}
	})
	group.Get("/info", func(ctx *fesgo.Context) {
		ctx.W.Write([]byte("get info"))
	})
	group.Post("/info", func(ctx *fesgo.Context) {
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
	engine.Run()
}
