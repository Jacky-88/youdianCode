package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"k8s-demo/config"
	"k8s-demo/controller"
	"k8s-demo/db"
	"k8s-demo/middle"
	"k8s-demo/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//初始化k8s clientset
	service.K8s.Init()
	//service.HelmConfig.Init()
	//初始化数据库
	db.Init()
	//初始化路由配置
	r := gin.Default()
	//跨域配置
	r.Use(middle.Cors())
	//jwt token验证
	//r.Use(middle.JWTAuth())
	//初始化路由
	controller.Router.InitApiRouter(r)

	//event任务,用于监听event并写入数据库,这里的传参是集群名，一定要与config中的集群名对齐
	//go func() {
	//	service.Event.WatchEventTask("TST-1")
	//}()
	//go func() {
	//	service.Event.WatchEventTask("TST-2")
	//}()

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

	//gin server启动
	srv := &http.Server{
		Addr:    config.ListenAddr,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//等待中断信号，优雅关闭所有server及DB
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//设置ctx超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//cancel用于释放ctx
	defer cancel()

	//关闭websocket
	if err := ws.Shutdown(ctx); err != nil {
		log.Fatal("Websocket关闭异常:", err)
	}
	log.Println("Websocket退出成功")

	//关闭gin server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Gin Server关闭异常:", err)
	}
	log.Println("Gin Server退出成功")
	//关闭db
	if err := db.Close(); err != nil {
		log.Fatal("DB关闭异常:", err)
	}
}