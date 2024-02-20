package main

import (
	"errors"
	"fmt"
	"github.com/dalefeng/fesgo"
	"github.com/dalefeng/fesgo/ferror"
	fesLog "github.com/dalefeng/fesgo/logger"
	"github.com/dalefeng/fesgo/render"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Name    string   `xml:"name" json:"name" validate:"required,max=5,min=2"`
	Address []string `xml:"address" json:"address"`
}

func main() {
	engine := fesgo.Default()
	group := engine.Group("user")
	group.Use(fesgo.Logging)

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
	group.Post("/string", func(c *fesgo.Context) {
		name := c.GetPostForm("name")
		ids, _ := c.GetPostFormMap("user")
		c.String(http.StatusOK, "string hello: %s, ids: %v", name, ids)
	})

	group.Post("/file", func(c *fesgo.Context) {
		file := c.FormFile("file")
		err := c.SaveUploadFile(file, "upload/"+file.Filename)
		if err != nil {
			log.Println(err)
			return
		}
		c.String(http.StatusOK, "success")
	})

	engine.Logger.SetLoggerPath("./log")
	engine.Logger.Formatter = &fesLog.TextFormatter{}
	engine.Logger.SetLevel(fesLog.LevelDebug)
	engine.Logger.LogFileSize = 1 << 10
	group.Post("/json", func(c *fesgo.Context) {
		var us []User
		err := c.BindJson(&us)
		if err != nil {
			log.Println(err)
			c.String(http.StatusOK, err.Error())
			return
		}
		c.Logger.Info("Info message", "name", "Feng")
		fesError := ferror.Default()
		fesError.Result(func(err *ferror.FesError) {
			c.Logger.Errorw("testErr", "err", err)
			c.String(http.StatusInternalServerError, err.Error())
			return
		})
		err = testErr(1)
		fesError.Put(err)
		err = testErr(2)
		fesError.Put(err)
		c.JSON(http.StatusOK, us)
	})

	fmt.Println("server run ...")
	engine.Run()
}

func testErr(p int) error {
	if p == 1 {
		return errors.New("testErr," + strconv.Itoa(p))
	}
	return nil
}
