package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// binding:"required" 修饰的字段，若接收为空值，则报错，是必须字段
type User struct {
	Name     string `json:"name" binding:"required"`
	Password int64  `json:"password" binding:"required"`
}

func Login1(c *gin.Context) {
	json := User{}
	c.BindJSON(&json)
	log.Printf("%v", &json)
	c.JSON(http.StatusOK, json)
}

func Login2(c *gin.Context) {
	json := make(map[string]interface{}) // 注意该结构接受的内容
	c.BindJSON(&json)
	log.Printf("%v", &json)
	c.JSON(http.StatusOK, json)
}

/* 获取查询参数 */
func main() {
	r := gin.Default()

	// ========================================================================
	// GET

	// 普通获取
	r.GET("/", func(c *gin.Context) {
		c.String(200, c.Query("wechat"))
	})

	// 默认值获取方式
	r.GET("/default", func(c *gin.Context) {
		c.String(200, c.DefaultQuery("wechat", "default"))
	})

	// 校验是否有传对应值
	r.GET("/getQuery", func(c *gin.Context) {
		value, ok := c.GetQuery("wechat")
		fmt.Println(ok)
		c.String(200, value)
	})

	// 获取参数 Array
	// ?media=100&media=200
	r.GET("/getQueryArray", func(c *gin.Context) {
		c.JSON(200, c.QueryArray("media"))
	})

	// 获取参数 Map
	// ?ids[a]=123&ids[b]=456&ids[c]=789
	r.GET("/getQueryMap", func(c *gin.Context) {
		c.JSON(200, c.QueryMap("ids"))
	})

	// ========================================================================
	// POST

	// 普通获取
	r.POST("/", func(c *gin.Context) {
		c.String(200, c.PostForm("wechat"))
	})

	// 默认值获取方式
	r.POST("/default", func(c *gin.Context) {
		c.String(200, c.DefaultPostForm("wechat", "default"))
	})

	// 校验是否有传对应值
	r.POST("/POSTQuery", func(c *gin.Context) {
		value, ok := c.GetPostForm("wechat")
		fmt.Print(ok)
		c.String(200, value)

	})

	// 获取参数 Array
	// 和 get 传参一致 只能获取 form-data 数据
	r.POST("/POSTQueryArray", func(c *gin.Context) {
		c.JSON(200, c.PostFormArray("media"))
	})

	// 获取参数 Map
	// ?ids[a]=123&ids[b]=456&ids[c]=789
	r.POST("/POSTQueryMap", func(c *gin.Context) {
		c.JSON(200, c.PostFormMap(""))
	})

	// 获取 json 参数（结构体）只能接收确定的参数，而且类型需要匹配
	/*
	   type User struct {
	       Name     string `json:"name"`
	       Password int64  `json:"password"`
	   }
	   json := &User{}
	   c.BindJSON(&json)
	*/
	r.POST("/Json1", Login1)

	// 获取 json 参数（接口）可以接收不确定的参数
	/*
	   json := make(map[string]interface{}) // 注意该结构接受的内容
	   c.BindJSON(&json)
	*/
	r.POST("/Json2", Login2)

	r.Run(":8080")

}
