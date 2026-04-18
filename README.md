- [English](README.md)
- [中文](README_cn.md)

## 📚 Introduction
gin-auto-router is an automatic routing component for the Golang Gin framework. It implements automatic registration of Gin routes, eliminating the need for manual addition and avoiding the complexities and omissions that come with manual route configuration. This allows developers to focus on business logic development while reducing the difficulty for beginners to get started with the Gin framework.

### 🧬 Project Features
This project supports all RESTful API request types, including: "post", "get", "put", "patch", "head", "options", "delete" and "any". It also supports multiple common naming conventions, such as: kebab-case (default), snake_case, camelCase, and PascalCase.

## 🎁 Changelog

### 🍺 v1.3.0
- **Important Updates:**
  1. Default naming convention changed to kebab-case
  2. Go language version upgraded to 1.26
  3. Gin framework version upgraded to v1.12.0
  4. Refactored route registration and binding methods
- **Compatibility Notes:**
  - Backward compatible, versions before v1.2.0 can continue to be used
  - If you were using the default binding method before, after upgrading to v1.3.0, you need to manually specify the naming convention:
    - Change `ginAutoRouter.Bind(engine)` to `ginAutoRouter.Bind(engine, "snake_case")`
    - Change `ginAutoRouter.BindGroup(routerGroup)` to `ginAutoRouter.BindGroup(routerGroup, "snake_case")`
  - If you were using a non-default binding method, you can continue using it after upgrading
  - No other code changes needed

### 🍺 v1.2.0
Rewrote route registration and binding methods while maintaining compatibility with v1.1.0. Added support for multiple naming conventions: snake_case (default), camelCase, and PascalCase.

### 🍺 v1.1.0
- Fixed a bug in route binding to ensure routes are correctly bound

### 🍺 v1.0.0
- Supports all RESTful API request types: post, get, put, patch, head, options, delete, any, where "any" binds to any request method
- Request URIs are automatically converted to lowercase snake_case format, e.g., ListGet → list_get, InfoPush → info_push

### 🍺 v0.2.0
Optimized registration and binding logic

## 📝 Quick Start

### Example Code [demo1 (Click to View)](examples/demo1)
- Assume the controller file is Article.go
- Method named Test or TestPost: Frontend uses POST to request "/article/test"
- Method named TestGet: Frontend uses GET to request "/article/test"
- Method named TestPut: Frontend uses PUT to request "/article/test"
- Method named TestPatch: Frontend uses PATCH to request "/article/test"
- Method named TestHead: Frontend uses HEAD to request "/article/test"
- Method named TestOptions: Frontend uses OPTIONS to request "/article/test"
- Method named TestDelete: Frontend uses DELETE to request "/article/test"
- Method named TestAny: Frontend uses any method to request "/article/test"

## 🛠️ Usage

### 🍐 Method 1
Use Gin's basic routing object to automatically generate access routes. Routes are automatically bound in the normal route binding mode.
- Example Code [demo1 (Click to View)](examples/demo1)

Main code is as follows:

- Entry main file
```sh
$ cat main.go
```

```go
package main

import "demo1/router"

func main() {
	// Load routes
	r := router.InitRouter()
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

- Route initialization file
```sh
# File in router directory init_router.go
$ cat router/init_router.go
```

```go
package router

import (
	_ "demo1/app/controller" // Must import controller package to register methods to be accessed
	ginAutoRouter "gitee.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Initialize routes
	r := gin.Default()
	// Auto-bind camelCase routes
	ginAutoRouter.Bind(r, "camelCase")
	return r
}
```

- Controller file
```sh
# File in controller directory article.go
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

### 🍊 Method 2
Route grouping mode: Use Gin's Group feature to implement route grouping, and apply middleware such as login verification to specific groups.
- Example Code [demo2 (Click to View)](examples/demo2)

Implementation effect: Login endpoint "/auth" can be accessed directly, but only logged-in users can access "/v1/article/list".
Most files are similar to Method 1, only route-related files have slight differences. For details, please refer to: Example Code [demo2 (Click to View)](examples/demo2)

Route-related file:
```sh
# File in router directory init_router.go
$ cat router/init_router.go
```

```go
package router

import (
	_ "demo2/app/controller/v1" // Must import controller package to register methods to be accessed
	"demo2/router/middleware/jwt"
	ginAutoRouter "gitee.com/dengshulei/gin-auto-router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Initialize routes
	r := gin.Default()
	// Enable v1 group
	v1Route := r.Group("/v1")
	// Load and use login verification middleware
	v1Route.Use(jwt.JWT())
	{
		// Bind Group routes (PascalCase)
		ginAutoRouter.BindGroup(v1Route, "PascalCase")
	}

	return r
}
```

- Except for the route binding code above, other code is basically the same as Method 1.