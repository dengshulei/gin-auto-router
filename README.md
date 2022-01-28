# 项目目标
#### 1. 路由自动注册，不用每次都手动添加
#### 2. 将精力放到具体业务逻辑开发上来
#### 3. 降低新手的入门难度

# 注意事项
#### 暂时只实现了post请求方式，这也是在生产中推荐使用的方式，其它方式可以参考代码自己实现，比较简单。


# 简要描述
#### 开发人员只需要编辑controller相关文件Article.go，就能实现轻松的访问“/Article/List”，类似于Nginx的路由方式




# 使用方法一
#### 示例代码[demo1](https://github.com/dengshulei/gin-auto-router/tree/master/examples/demo1)

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
	//绑定基本路由，访问路径：/Article/List
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
```

# 使用方法二：

#### 示例代码[demo2](https://github.com/dengshulei/gin-auto-router/tree/master/examples/demo2)

#### 利用gin的分组路由Group实现路由分组，并对指定的分组使用登录验证等中间件。
#### 实现的效果是：登录接口“/auth”可以直接访问，只有登录成功的才可以访问“/v1/Article/List”

#### 大多的文件都是与方法一类似，只有路由相关文件有细微的区别，详细的情况请查看示例代码 [demo2](https://github.com/dengshulei/gin-auto-router/tree/master/examples/demo2)

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
		//绑定Group路由，访问路径：/v1/Article/List
		ginAutoRouter.BindGroup(v1Route)
	}

	return r
}

```

