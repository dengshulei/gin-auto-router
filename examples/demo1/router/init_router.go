package router

import (
	_ "demo1/app/controller" // Must import this controller package to register methods to be accessed
	gin_auto_router "gitee.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes router
// Naming methods supported: snake_case (default), camelCase, PascalCase
func InitRouter() *gin.Engine {
	// Initialize router
	engine := gin.Default()

	// Auto bind routes
	gin_auto_router.Bind(engine, "kebab-case")
	//gin_auto_router.Bind(engine, "snake_case")
	//gin_auto_router.Bind(engine, "camelCase")
	//gin_auto_router.Bind(engine, "PascalCase")
	//gin_auto_router.Bind(engine)

	return engine
}
