package main

import (
	"fmt"
	"github.com/starbuling-l/StarBlog/pkg/setting"
	"github.com/starbuling-l/StarBlog/routers"
	"net/http"
)

/**
测试
curl 127.0.0.1:9000/test
{"message":"test"}
*/

func main() {
	r := routers.InitRouter()
	s := &http.Server{
		Handler:           r,
		Addr:              fmt.Sprintf(":%d", setting.HTTPPort),
		ReadHeaderTimeout: setting.ReadTimeOut,
		WriteTimeout:      setting.WriteTimeOut,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
}