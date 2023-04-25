package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"kubea-demo/config"
	"kubea-demo/controller"
	"kubea-demo/db"
	"kubea-demo/middle"
	"kubea-demo/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//初始化gin对象
	r := gin.Default()
	//初始化数据库
	db.Init()
	//初始化k8s client
	service.K8s.Init()
	//跨域中间件注册
	r.Use(middle.Cors())
	//r.Use(middle.JWTAuth())
	//初始化路由
	controller.Router.InitApiRouter(r)
	//启动task
	go func() {
		service.Event.WatchEventTask("TST-1")
	}()
	//websocket 启动
	wsHandler := http.NewServeMux()
	wsHandler.HandleFunc("/ws", service.Terminal.WsHandler)
	ws := &http.Server{
		Addr:    config.WsAddr,
		Handler: wsHandler,
	}
	go func() {
		if err := ws.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	//启动gin server
	srv := &http.Server{
		Addr: config.ListenAddr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen: %s\n", err)
		}
	}()
	//优雅关闭server
	//声明一个系统信号的channel，并监听他，如果没有信号，就一直阻塞，如果有，就继续执行
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	//设置ctx超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	//cancel用于释放ctx
	defer cancel()
	//关闭websocket
	if err := ws.Shutdown(ctx); err != nil {
		log.Fatal("Websocket关闭异常:", err)
	}
	log.Println("Websocket退出成功")
	//关闭gin
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Gin Server关闭异常: ", err)
	}
	logger.Info("Gin Server退出成功")
	//关闭db
	if err := db.Close(); err != nil {
		log.Fatal("DB关闭异常:", err)
	}
}
