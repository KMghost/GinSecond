package main

import "github.com/gin-gonic/gin"

type User1 struct {
	ID   int
	Name string
}
type User2 struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

func main() {
	r := gin.Default()

	// 普通写法
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"wechat": "flysnow_org", "blog": "www.flysnow.org"})
	})

	// ========================================================================
	// 自定义 struct
	// type User1 struct {ID int; Name string}
	r.GET("/xml1", func(c *gin.Context) {
		c.XML(200, User1{ID: 123, Name: "张三"})
	})

	// ========================================================================
	// 自定义节点名称
	// type User2 struct {ID int `xml:"id"`; Name string `xml:"name"`}
	r.GET("/xml2", func(c *gin.Context) {
		c.XML(200, User2{ID: 123, Name: "张三"})
	})

	// ========================================================================
	// xml 数组
	r.GET("/xml3", func(c *gin.Context) {
		allUsers := []User2{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}
		c.XML(200, gin.H{"user": allUsers})
	})

	r.Run(":8080")
}
