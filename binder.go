package gin_auto_router

import (
	"github.com/gin-gonic/gin"
)

// Bind 绑定基本路由，外部可以直接使用
func Bind(e *gin.Engine, args ...string) {
	r := e.Group("/")
	BindGroup(r, GetNom(args...))
}

// BindGroup 绑定路由组，外部可以直接使用
func BindGroup(g *gin.RouterGroup, args ...string) {
	for _, r := range Routes {
		r.Ext.Nom = GetNom(args...)
		// 执行路由绑定
		r.Bind(g)
	}
}
