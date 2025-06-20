package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	staticDir := "static"

	// 读取 static 目录下的所有子目录
	files, err := os.ReadDir(staticDir)
	if err != nil {
		panic(fmt.Sprintf("无法读取 static 目录: %v", err))
	}

	for _, file := range files {
		if file.IsDir() {
			subDir := file.Name()
			// 注册路由：例如 /mark -> static/mark/dist
			r.StaticFS("/"+subDir, http.Dir(staticDir+"/"+subDir+"/dist"))
			fmt.Printf("已注册路由: /%s -> %s/%s/dist\n", subDir, staticDir, subDir)
		}
	}

	// 启动服务
	fmt.Println("服务启动中，监听 :8081")
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
