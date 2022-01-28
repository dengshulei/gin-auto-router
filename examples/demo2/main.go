package main

import (
	"demo2/router"
)

func main() {
	//加载路由
	r := router.InitRouter()
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}