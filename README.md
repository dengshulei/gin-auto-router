##  📚 简介
gin-auto-router 是 Golang 语言 Gin 框架的自动路由组件。实现了 Gin 路由的自动注册，不用手动添加，避免了手动添加路由的一系列问题的出现。 将开发人员的精力放到具体业务逻辑开发上来。 同时降低了新手的入门 Golang 下 Gin 框架开发的难度。

### 🧬 项目特点
本项目支持所有的 RESTful API 请求类型，包括："post", "get", "put", "patch", "head", "options", "delete" 和 "any"等。本项目支持常见的命名方法(nomenclature),比如: snake_case (下划线命名法，默认)、camelCase(驼峰命名法)和 PascalCase (帕斯卡命名法)。

## 🎁 升级记录
### 🍺 v1.2.0.
在保持与 v1.1.0 兼容的基础上，重写了路由注册与绑定的方法。新增支持多种常见的命名方法( nomenclature ),比如: snake_case (下划线命名法，默认)、camelCase(驼峰命名法)和 PascalCase (帕斯卡命名法)。
### 🍺 v1.1.0
从 v1.0.0 开始，支持所有的 RESTful API 请求类型，支持的类型分别是：post/get/put/patch/head/options/delete/any，其中 any 将绑定为为任意请求方式。
### 🍺 v1.1.0
从 v1.0.0 开始，请求 URI 自动为小写下划线方式，即snake_case (下划线命名法) 。例如：ListGet => list_get, InfoPush => info_push
### 🍺 v0.2.0
优化了注册和绑定的逻

## 📝 简单说明
### 示例代码 [demo1(点击直达)](/examples/demo1)
- 假定控制器文件为 Article.go
- 方法名为 Test 或 TestPost：前端使用 post 方式请求 “/article/test”
- 方法名为 TestGet：前端使用 get 方式请求 “/article/test”
- 方法名为 TestPut：前端使用 put 方式请求 “/article/test”
- 方法名为 TestPatch：前端使用 patch 方式请求 “/article/test”
- 方法名为 TestHead：前端使用 head 方式请求 “/article/test”
- 方法名为 TestOptions：前端使用 options 方式请求 “/article/test”
- 方法名为 TestDelete：前端使用 delete 方式请求 “/article/test”
- 方法名为 TestAny：前端使用 任意 方式请求 “/article/test”

## 🛠️ 使用方法
### 🍐 使用方法 1
利用gin基本的路由对象，自动生成访问路由。即在路由绑定的普通模式下自动绑定油路。
- 示例代码 [demo1 (点击直达)](/examples/demo1)

主要代码如下：

- 入口主文件
```sh
$ cat main.go
```

```go
package main
import "demo1/router"

func main() {
	//加载路由
	r := router.InitRouter()
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

- 自动生成路由的主文件
```sh
# router下的文件InitRouter.go
$ cat InitRouter.go
```

```go
package router
import (
	_ "demo1/app/controller"   //一定要导入这个Controller包，用来注册需要访问的方法
	ginAutoRouter "github.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)
func InitRouter() *gin.Engine  {
	//初始化路由
	r := gin.Default()
	//自动绑定 驼峰命名法 路由
	ginAutoRouter.Bind(r, "camelCase")
	return r
}
```

- 控制器文件
```sh
# controller下的文件Article.go
$ cat Article.go
```

```go
package controller

import (
	ginAutoRouter "github.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	ginAutoRouter.Register(&Article{})
}

type Article struct {}

func (api *Article) ListGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "ok",
		"data": "Article:List",
	})
}

func (api *Article) InfoGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "ok",
		"data": "Article:InfoGet",
	})
}

func (api *Article) InfoPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "ok",
		"data": "Article:InfoPost",
	})
}

func (api *Article) InfoDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "ok",
		"data": "Article:InfoDelete",
	})
}
func (api *Article) InfoPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "ok",
		"data": "Article:InfoPut",
	})
}
```

### 🍊 使用方法 2
路由分组模式，即：利用 Gin 的分组路由 Group 实现路由分组，并对指定的分组使用登录验证等中间件。
- 示例代码 [demo2(点击直达)](/examples/demo2)

实现的效果是：登录接口“/auth”可以直接访问，但是只有登录成功的才可以访问“/v1/article/list”。
- 大多的文件都是与方法一类似，只有路由相关文件有细微的区别，详细的情况请查看：示例代码 [demo2(点击直达)](/examples/demo2) 

路由相关文件
```sh
$ cat InitRouter.go
```

```go
package router

import (
	_ "demo2/app/controller/v1" //一定要导入这个Controller包，用来注册需要访问的方法
	"demo2/router/middleware/jwt"
	ginAutoRouter "github.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine  {
	//初始化路由
	r := gin.Default()
	//开启v1分组
	v1Route := r.Group("/v1")
	//加载并使用登录验证中间件
	v1Route.Use(jwt.JWT())
	{
		//绑定Group路由 (帕斯卡命名法)
		ginAutoRouter.BindGroup(v1Route, "PascalCase")
	}

	return r
}

```
- 除了上面的绑定路由的代码，其它代码都与“使用方法 1”基本一致。