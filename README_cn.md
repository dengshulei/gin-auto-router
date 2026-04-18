- [English](README.md)
- [中文](README_cn.md)

##  📚 简介
gin-auto-router 是 Golang 语言 Gin 框架的自动路由组件。它实现了 Gin 路由的自动注册，无需手动添加，避免了手动配置路由的繁琐与遗漏问题。让开发者专注于业务逻辑开发，同时降低了新手入门 Gin 框架的难度。

### 🧬 项目特点
本项目支持所有 RESTful API 请求类型，包括："post"、"get"、"put"、"patch"、"head"、"options"、"delete" 和 "any" 等。同时支持多种常见的命名规范，如：kebab-case（短横线命名法，默认）、snake_case（下划线命名法）、camelCase（驼峰命名法）和 PascalCase（帕斯卡命名法）。

## 🎁 更新日志

### 🍺 v1.3.0
- **重要更新：**
  1. 默认命名规范调整为 kebab-case（短横线命名法）
  2. Go 语言版本升级为 1.26
  3. Gin 框架版本升级为 v1.12.0
  4. 重构了路由注册与绑定方法
- **兼容提示：**
  - 保持向后兼容，v1.2.0 之前的版本可继续使用
  - 若之前使用默认绑定方式，升级到 v1.3.0 后需手动指定命名规范：
    - 将 `ginAutoRouter.Bind(engine)` 改为 `ginAutoRouter.Bind(engine, "snake_case")`
    - 将 `ginAutoRouter.BindGroup(routerGroup)` 改为 `ginAutoRouter.BindGroup(routerGroup, "snake_case")`
  - 若之前已使用非默认方式，升级后可直接继续使用
  - 其他代码无需改动

### 🍺 v1.2.0
在保持与 v1.1.0 兼容的基础上，重写了路由注册与绑定方法。新增支持多种命名规范：snake_case（下划线命名法，默认）、camelCase（驼峰命名法）和 PascalCase（帕斯卡命名法）。

### 🍺 v1.1.0
- 修正绑定路由时的 bug，使其能够正确绑定路由

### 🍺 v1.0.0
- 支持所有 RESTful API 请求类型：post、get、put、patch、head、options、delete、any，其中 "any" 绑定为任意请求方式
- 请求 URI 自动转换为小写下划线格式（snake_case），如：ListGet → list_get、InfoPush → info_push
- 
### 🍺 v0.2.0
优化了注册与绑定的逻辑

## 📝 快速入门

### 示例代码 [demo1（点击直达）](/examples/demo1)
- 假设控制器文件为 Article.go
- 方法名为 Test 或 TestPost：前端使用 POST 请求地址 "/article/test"
- 方法名为 TestGet：前端使用 GET 请求地址 "/article/test"
- 方法名为 TestPut：前端使用 PUT 请求地址 "/article/test"
- 方法名为 TestPatch：前端使用 PATCH 请求地址 "/article/test"
- 方法名为 TestHead：前端使用 HEAD 请求地址 "/article/test"
- 方法名为 TestOptions：前端使用 OPTIONS 请求地址 "/article/test"
- 方法名为 TestDelete：前端使用 DELETE 请求地址 "/article/test"
- 方法名为 TestAny：前端使用任意方式请求地址 "/article/test"

## 🛠️ 使用方法

### 🍐 方法一
利用 Gin 基本的路由对象，自动生成访问路由。即在路由绑定的普通模式下自动绑定路由。
- 示例代码 [demo1（点击直达）](/examples/demo1)

主要代码如下：

- 入口主文件
```sh
$ cat main.go
```

```go
package main

import "demo1/router"

func main() {
	// 加载路由
	r := router.InitRouter()
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

- 路由初始化文件
```sh
# router 目录下的 init_router.go
$ cat router/init_router.go
```

```go
package router

import (
	_ "demo1/app/controller" // 必须导入 controller 包，用于注册需要访问的方法
	ginAutoRouter "gitee.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 初始化路由
	r := gin.Default()
	// 自动绑定 camelCase 路由
	ginAutoRouter.Bind(r, "camelCase")
	return r
}
```

- 控制器文件
```sh
# controller 目录下的 article.go
$ cat app/controller/article.go
```

```go
package controller

import (
	ginAutoRouter "gitee.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	ginAutoRouter.Register(&Article{})
}

type Article struct{}

func (api *Article) ListGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:List",
	})
}

func (api *Article) InfoGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:InfoGet",
	})
}

func (api *Article) InfoPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:InfoPost",
	})
}

func (api *Article) InfoDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:InfoDelete",
	})
}

func (api *Article) InfoPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": "Article:InfoPut",
	})
}
```

### 🍊 方法二
路由分组模式：利用 Gin 的分组路由（Group）实现路由分组，并对指定分组使用登录验证等中间件。
- 示例代码 [demo2（点击直达）](/examples/demo2)

实现效果：登录接口 "/auth" 可直接访问，但只有登录成功后才能访问 "/v1/article/list"。
大部分文件与方法一类似，仅路由相关文件有细微区别，详细请查看：示例代码 [demo2（点击直达）](/examples/demo2)

路由相关文件：
```sh
# router 目录下的 init_router.go
$ cat router/init_router.go
```

```go
package router

import (
	_ "demo2/app/controller/v1" // 必须导入 controller 包，用于注册需要访问的方法
	"demo2/router/middleware/jwt"
	ginAutoRouter "gitee.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 初始化路由
	r := gin.Default()
	// 启用 v1 分组
	v1Route := r.Group("/v1")
	// 加载并使用登录验证中间件
	v1Route.Use(jwt.JWT())
	{
		// 绑定 Group 路由（PascalCase）
		ginAutoRouter.BindGroup(v1Route, "PascalCase")
	}

	return r
}
```

- 除上述路由绑定代码外，其他代码与方法一基本一致。