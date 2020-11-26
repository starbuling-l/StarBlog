package main

import (
	"context"
	"fmt"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"github.com/starbuling-l/StarBlog/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 测试 curl 127.0.0.1:9000/test
//{"message":"test"}

//优雅重启
//每次更新发布、或者修改配置文件等 给该进程发送SIGTERM 信号 ，而不需要强制结束应用
//问题 endless 热更新是采取创建子进程后，将原进程退出的方式，这点不符合守护进程的要求
//func main() {
//	endless.DefaultWriteTimeOut = setting.WriteTimeOut
//	endless.DefaultReadTimeOut = setting.ReadTimeOut
//	endless.DefaultMaxHeaderBytes = 1 << 20
//
//	Addr := fmt.Sprintf(":%d", setting.HTTPPort)
//	router := routers.InitRouter()
//	server := endless.NewServer(Addr, router)
//	server.BeforeBegin = func(add string) {
//		log.Printf("current pid is %d", syscall.Getpid())
//	}
///*
//	s := &http.Server{
//		Handler:           r,
//		Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
//		ReadHeaderTimeout: setting.ReadTimeOut,
//		WriteTimeout:      setting.WriteTimeOut,
//		MaxHeaderBytes:    1 << 20,
//	}*/
//	if err := server.ListenAndServe(); err != nil {
//		log.Printf("Server err:%v", err)
//	}
//}

func main() {

	Addr := fmt.Sprintf(":%d", setting.HTTPPort)
	router := routers.InitRouter()

	server := &http.Server{
		Handler:           router,
		Addr:              Addr,
		ReadHeaderTimeout: setting.ReadTimeOut,
		WriteTimeout:      setting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Server err:%v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("shutdown server ....")
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Hour)
	defer cancel()

	if err := server.Shutdown(cxt); err != nil {
		log.Fatal("shutdown err :", err)
	}
	log.Println("server exiting")
}