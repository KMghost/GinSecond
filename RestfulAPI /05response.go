package main

import "github.com/gin-gonic/gin"

/*
   json 渲染输出
*/
type user1 struct {
	ID   int
	Name string
	Age  int
}

type user2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()

	// json 响应的简单使用
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world"})
	})

	// struct 转 json
	// type user struct {ID   int; Name string; Age  int}
	r.GET("/users/1", func(c *gin.Context) {
		c.JSON(200, user1{ID: 123, Name: "张三", Age: 20})
	})

	// 自定义 json 字段的名字
	// type user2 struct {ID   int `json:"id"`; Name string `json:"name"`; Age  int `json:"age"`}
	r.GET("/users/2", func(c *gin.Context) {
		c.JSON(200, user2{ID: 123, Name: "张三", Age: 20})
	})

	// JSONP (主要为了给前端防止跨域)
	// url: /jsonp?callback=aa,           response: aa({"id":123,"name":"张三","age":20});
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, user2{ID: 123, Name: "张三", Age: 20})
	})

	// json数组
	// 直接传一个 数组到 Json 中，IndentedJSON 自动格式化 json
	allUsers := []user2{{ID: 123, Name: "张三", Age: 20}, {ID: 456, Name: "李四", Age: 25}}
	r.GET("/users/3", func(c *gin.Context) {
		c.IndentedJSON(200, allUsers)
	})

	// PureJSON
	// 对于 Json 字符串中特殊的字符串，比如 '<' ，Gin 默认是转义的，比如变成 '\u003c'，需要保持原来的字符，不进行转使用 PureJSON
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "<b>Hello, world!</b>",
		})
	})
	r.GET("/pureJson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"message": "<b>Hello, world!</b>",
		})
	})

	// AsciiJSON
	// 把非 Ascii 字符串转为 unicode 编码，使用 AsciiJSON
	r.GET("/asciiJSON", func(c *gin.Context) {
		c.AsciiJSON(200, gin.H{"message": "hello 飞雪无情"})
	})

	r.Run(":8080")
}
