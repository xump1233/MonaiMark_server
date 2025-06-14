package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 设置静态文件服务：将 dist 目录映射为网站根路径 /
	r.StaticFS("/", http.Dir("static/dist"))

	// 启动服务
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
