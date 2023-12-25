package main

import (
	"github.com/dalefeng/fesgo"
	"net/http"
)

func main() {
	engine := fesgo.NewEngine("8111")
	g := engine.Group("user")
	g.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello fes"))
	})
	engine.Run()
}
