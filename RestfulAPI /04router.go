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

	r.Run()
}
