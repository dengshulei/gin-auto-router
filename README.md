# Go语言的Gin框架自动路由组件
# 1 项目目标
### 1.1 Gin的路由自动注册，不用每次都手动添加
### 1.2 将精力放到具体业务逻辑开发上来
### 1.3 降低新手的入门难度

# 2 特别说明
### 2.1 从 v1.0.0 开始，支持所有的 RESTful API 请求类型
### 2.2 从 v1.0.0 开始，请求 URI 自动为小写下划线方式

# 3 简要描述
### 3.1 假定控制器目录为 controller，具体可参见：示例代码[demo1](/examples/demo1)
### 3.2 假定控制器文件为 Article.go，方法名为 List
### 3.3 前端使用 post 方式访问“/article/list”，即可访问到上述代码
### 3.4 从 v1.0.0 开始，控制器文件 Article.go 下的方法名为 TestGet 则必须用 get 方式请求“/article/list”
### 3.2 从 v1.0.0 开始，控制器文件 Article.go 下的方法名为 TestGet 则必须用get方式请求“/article/list”
### 3.2 从 v1.0.0 开始，控制器文件 Article.go 下的方法名为 TestGet 则必须用get方式请求“/article/list”
### 3.2 从 v1.0.0 开始，控制器文件 Article.go 下的方法名为 TestGet 则必须用get方式请求“/article/list”



# 使用方法一
#### 示例代码[demo1](/examples/demo1)

#### 利用gin基本的路由对象，自动生成访问路由，主要代码如下：

### 入口主文件
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

### 自动生成路由的主文件
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
	//绑定基本路由，访问路径：/article/list
	ginAutoRouter.Bind(r)
	return r
}
```

### 控制器文件
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

func (api *Article) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "ok",
		"data": "Article:List",
	})
}

func (api *Article) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "ok",
		"data": "Article:Test",
	})
}
func (api *Article) TestGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": "ok",
		"data": "Article:Test",
	})
}
```

# 使用方法二：

#### 示例代码[demo2](/examples/demo2)

#### 利用gin的分组路由Group实现路由分组，并对指定的分组使用登录验证等中间件。
#### 实现的效果是：登录接口“/auth”可以直接访问，只有登录成功的才可以访问“/v1/Article/List”

#### 大多的文件都是与方法一类似，只有路由相关文件有细微的区别，详细的情况请查看示例代码 [demo2](/examples/demo2)

### 路由相关文件
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
		//绑定Group路由，访问路径：/v1/article/list
		ginAutoRouter.BindGroup(v1Route)
	}

	return r
}

```

