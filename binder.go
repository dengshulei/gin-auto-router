package gin_auto_router

import (
	"github.com/gin-gonic/gin"
)

// Bind binds all registered routes to the root group "/".
// This is a convenience method when you don't need custom route grouping.
//
// Parameters:
//
//	engine - The gin.Engine instance to bind routes to
//	naming - (Optional) Route naming convention.
//	        Supported: kebab-case (default), snake_case, camelCase, PascalCase
//
// Example:
//
//	import (
//	// Import controller package to register methods
//	_ "demo1/app/controller"
//	ginAutoRouter "gitee.com/dengshulei/gin-auto-router"
//	"github.com/gin-gonic/gin"
//
//	)
//
//	func InitRouter() *gin.Engine {
//	    engine := gin.Default()
//	    ginAutoRouter.Bind(engine)                    // Use default kebab-case
//	    //ginAutoRouter.Bind(engine, "camelCase")        // Use camelCase
//	    return engine
//	}
func Bind(engine *gin.Engine, naming ...string) {
	routerGroup := engine.Group("/")
	BindGroup(routerGroup, NormalizeNamingConvention(naming...))
}

// BindGroup binds all registered routes to a custom router group.
// Use this when you need custom route group paths or want to add middleware.
//
// Parameters:
//
//	routerGroup - The gin.RouterGroup to bind routes to
//	naming - (Optional) Route naming convention.
//	        Supported: kebab-case (default), snake_case, camelCase, PascalCase
//
// Example:
//
//	import (
//	// Import controller package to register methods
//	_ "demo2/app/controller/v1"
//	"demo2/router/middleware/jwt"
//	ginAutoRouter "gitee.com/dengshulei/gin-auto-router"
//	"github.com/gin-gonic/gin"
//
//	)
//
//	func InitRouter() *gin.Engine {
//	    engine := gin.Default()
//
//	    apiGroup := engine.Group("/api")
//	    apiGroup.Use(middleware.Auth())  // Add middleware
//	    ginAutoRouter.BindGroup(apiGroup, "PascalCase")
//
//	    return engine
//	}
func BindGroup(routerGroup *gin.RouterGroup, naming ...string) {
	nom := NormalizeNamingConvention(naming...)
	// Iterate through all registered routes
	for _, route := range Routes {
		route.Ext.NamingConvention = nom
		// Perform route binding
		route.Bind(routerGroup)
	}
}
