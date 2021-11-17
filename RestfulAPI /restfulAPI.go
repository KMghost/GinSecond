package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/users", func(context *gin.Context) {
		//创建一个用户
	})
	r.DELETE("/usrs/123", func(context *gin.Context) {
		//删除ID为123的用户
	})
	r.PUT("/usrs/123", func(context *gin.Context) {
		//更新ID为123的用户
	})

	r.PATCH("/usrs/123", func(context *gin.Context) {
		//更新ID为123用户的部分信息
	})

	r.Any("/usr/123", func(context *gin.Context) {
		// 使用所有的方法操作
	})
	Handle(r, []string{"GET", "POST"}, "/", func(c *gin.Context) {
		//同时注册GET、POST请求方法
	})
}

// 自定义路由写法，不推荐使用，破坏了 resultful 规范中的 HTTP Method 的约束
func Handle(r *gin.Engine, httpMethods []string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	var routes gin.IRoutes
	for _, httpMethod := range httpMethods {
		routes = r.Handle(httpMethod, relativePath, handlers...)
	}
	return routes
}
