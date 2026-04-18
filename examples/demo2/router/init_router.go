package router

import (
	_ "demo2/app/controller/v1" // Must import this controller package to register methods to be accessed
	"demo2/router/middleware/jwt"
	ginAutoRouter "gitee.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

// InitRouter initializes router
// Naming methods supported: snake_case (default), camelCase, PascalCase
func InitRouter() *gin.Engine {
	// Initialize router
	engine := gin.Default()
	// Enable api group
	routerGroup := engine.Group("/api")
	// Load and use login verification middleware
	routerGroup.Use(jwt.JWT())
	{
		// Bind Group routes

		ginAutoRouter.BindGroup(routerGroup, "kebab-case")
		//ginAutoRouter.BindGroup(routerGroup, "snake_case")
		//ginAutoRouter.BindGroup(routerGroup, "camelCase")
		//ginAutoRouter.BindGroup(routerGroup, "PascalCase")
		//ginAutoRouter.BindGroup(routerGroup)
	}

	return engine
}
