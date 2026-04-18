package main

import (
	"demo2/router"
)

func main() {
	// Load router
	engine := router.InitRouter()
	_ = engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
