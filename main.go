package main

import (
	"github.com/dalefeng/fesgo"
)

func main() {
	engine := fesgo.NewEngine("8111")
	group := engine.Group("user")
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
	engine.Run()
}
