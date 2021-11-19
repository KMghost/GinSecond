package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ========================================================================
	// 分组路由 （避免重复的输入前面的url）

	v1Group := r.Group("/v1")
	{
		v1Group.GET("/users", func(c *gin.Context) {
			c.String(200, "/v1/users")
		})
		v1Group.GET("/products", func(c *gin.Context) {
			c.String(200, "/v1/products")
		})
	}

	// ========================================================================
	// 路由中间件
	v2Group := r.Group("/v2", func(c *gin.Context) {
		fmt.Println("中间件")
	})
	{
		v2Group.GET("/users", func(c *gin.Context) {
			c.String(200, "/v2/users")
		})
		v2Group.GET("/products", func(c *gin.Context) {
			c.String(200, "/v2/products")
		})
	}

	// ========================================================================
	// 分组路由嵌套
	v1AdminGroup := v1Group.Group("/admin")
	{
		v1AdminGroup.GET("/users", func(c *gin.Context) {
			c.String(200, "/v1/admin/users")
		})
		v1AdminGroup.GET("/manager", func(c *gin.Context) {
			c.String(200, "/v1/admin/manager")
		})
	}

	r.Run()
}
