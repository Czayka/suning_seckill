package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"suning/mongo"
	"time"
	"github.com/gin-gonic/gin"
)
type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
func (g *Gin) Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
	return
}

func StartServer(){
	gin.SetMode("debug")
	engine := gin.New()

	// 性能分析 - 正式环境不要使用！！！
	// pprof.Register(engine)

	// 设置路由

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := Gin{Ctx: c}
		utilGin.Response(200, "pong", nil)
	})
	engine.GET("/token/:id", func(c *gin.Context) {
		id := c.Param("id")
		dfpToken := mongo.DfpToken{}
		dfpToken.CreateTime = time.Now().Unix()
		dfpToken.Token = id
		dfpToken.Insert()
		utilGin := Gin{Ctx: c}
		utilGin.Response(200, "pong", nil)
	})
	server := &http.Server{
		Addr:         ":8888",
		Handler:      engine,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|            go-gin-api             |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port:8888" + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
