package main

import "github.com/gin-gonic/gin"

/* 路由传参 */
// 冒号路由匹配
// func main() {
//     r := gin.Default()
//
//     r.GET("/users/:id", func(c *gin.Context) {
//         id := c.Param("id")    // 获取定义的路由参数的值
//         c.String(200, "The user id is  %s", id)
//     })
//     r.Run(":8080")
// }

// 星号路由匹配
func main() {
	r := gin.Default()

	r.GET("/users/*id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "The user id is  %s", id)
	})
	r.Run(":8080")
}
