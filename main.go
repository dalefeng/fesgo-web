package main

import (
	"errors"
	"fmt"
	"github.com/dalefeng/fesgo"
	"github.com/dalefeng/fesgo-blog/service"
	"github.com/dalefeng/fesgo/fpool"
	fesLog "github.com/dalefeng/fesgo/logger"
	"github.com/dalefeng/fesgo/token"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Name    string   `xml:"name" json:"name" validate:"required,max=5,min=2"`
	Address []string `xml:"address" json:"address"`
}

func main() {
	engine := fesgo.Default()
	//config.LoadConfig()
	engine.RegisterErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *Response:
			return http.StatusOK, e
		default:
			return http.StatusInternalServerError, 500
		}
	})
	group := engine.Group("user")
	//auth := &fesgo.Account{
	//	Users: map[string]string{
	//		"feng": "123",
	//	},
	//}
	//j := token.JwtHandler{
	//	Secret: []byte("123456"),
	//}
	group.Use(fesgo.Logging)
	//group.Use(j.AuthInterceptor)
	//group.Use(fesgo.Logging, auth.BasicAuth)
	group.Use(func(next fesgo.HandlerFunc) fesgo.HandlerFunc {
		return func(c *fesgo.Context) {
			fmt.Println("pre handler1")
			next(c)
			fmt.Println("post handler")
		}
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

	group.Get("/file", func(c *fesgo.Context) {
		c.FileAttachment("tmp/main.exe", "man.exe")
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
	group.Post("/json", func(c *fesgo.Context) {
		var us []User
		err := c.BindJson(&us)
		if err != nil {
			log.Println(err)
			c.String(http.StatusOK, err.Error())
			return
		}
		c.Logger.Info("Info message", "name", "Feng")
		err = Error("账号错误")
		c.HandleError(http.StatusOK, nil, err)
	})

	pool, err := fpool.NewPool(2)
	if err != nil {
		panic(err)
	}
	group.Post("/pool", func(c *fesgo.Context) {
		g := sync.WaitGroup{}
		count := 3
		g.Add(count)
		start := time.Now()
		for i := 0; i < count; i++ {
			func(index int) {
				pool.Submit(func() {
					defer g.Done()
					c.Logger.Infow("pool", "index", index)
					time.Sleep(0 * time.Second)
				})
			}(i)
		}
		g.Wait()
		end := time.Now()
		c.Logger.Infow("time", "cost", end.Sub(start))
		c.String(http.StatusOK, "pool")
	})

	group.Get("/login", func(c *fesgo.Context) {
		jwt := &token.JwtHandler{}
		jwt.Secret = []byte("123456")
		jwt.SendCookie = true
		jwt.Expire = 10 * time.Minute
		jwt.RefreshExpire = 20 * time.Minute
		jwt.Authenticator = func(c *fesgo.Context) (map[string]any, error) {
			data := make(map[string]any)
			data["userId"] = 1
			return data, nil
		}
		tokenResp, err := jwt.LoginHandler(c)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, tokenResp)

	})

	//
	group.Get("/refresh", func(c *fesgo.Context) {
		jwt := &token.JwtHandler{}
		jwt.Secret = []byte("123456")
		jwt.SendCookie = true
		jwt.Expire = 10 * time.Minute
		jwt.RefreshKey = "refresh-token"
		jwt.Authenticator = func(c *fesgo.Context) (map[string]any, error) {
			data := make(map[string]any)
			data["userId"] = 1
			return data, nil
		}
		c.Set("refresh-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEyNTg4MDcsImlhdCI6MTcxMTI1NzYwNywidXNlcklkIjoxfQ.YN6I98xWFwVZAnbfZ4n4Oe2OTScd3KGIG18J3rBYcGw")
		tokenResp, err := jwt.RefreshHandler(c)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, tokenResp)

	})

	group.Get("/test", func(c *fesgo.Context) {
		service.Select()
	})

	fmt.Println("server run ...")
	engine.Run(":8111")
}

func testErr(p int) error {
	if p == 1 {
		return errors.New("testErr," + strconv.Itoa(p))
	}
	return nil
}

type Response struct {
	Code int
	Msg  string
	Data any
}

func (r *Response) Error() string {
	return r.Msg
}

func Success(data any) *Response {
	return &Response{Code: 0, Msg: "success", Data: data}
}

func Error(msg string) *Response {
	return &Response{Code: 500, Msg: msg, Data: nil}
}
