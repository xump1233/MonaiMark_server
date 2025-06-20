package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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

	// 保存所有子应用名
	var subApps []string

	for _, file := range files {
		if file.IsDir() {
			subDir := file.Name()
			subApps = append(subApps, subDir)

			// 注册静态资源路由，例如 /mark -> static/mark/dist
			r.StaticFS("/"+subDir, http.Dir(staticDir+"/"+subDir+"/dist"))
			fmt.Printf("✅ 已注册路由: /%s -> %s/%s/dist\n", subDir, staticDir, subDir)
		}
	}

	// NoRoute fallback 处理，用于支持前端路由
	r.NoRoute(func(c *gin.Context) {
		requestPath := c.Request.URL.Path

		// 遍历所有子应用前缀
		for _, app := range subApps {
			prefix := "/" + app

			if strings.HasPrefix(requestPath, prefix) {
				// 如果请求路径以子应用前缀开头，返回其 index.html
				indexPath := fmt.Sprintf("%s/%s/dist/index.html", staticDir, app)
				c.File(indexPath)
				return
			}
		}

		// 非前端应用路径，返回 404
		c.String(http.StatusNotFound, "404 Not Found")
	})

	// 启动服务
	fmt.Println(" 服务启动中，监听 :8081")
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}

