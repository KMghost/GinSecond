package main

import "github.com/gin-gonic/gin"

/*
   浏览器都遵循同源策略，也就是说位于www.flysnow.org下的网页是无法访问非www.flysnow.org下的数据的，比如我们常见的AJAX跨域问题。

   要解决跨域问题的办法有CORS、代理和JSONP，这里结合Gin，主要介绍JSONP模式

   1、JSONP原理
       JSONP可以跨域，主要是利用了<script>跨域的能力，因为这个标签我们可以引用任何域名下的JS文件。
       既然是这样，我们就可以利用这个能力，在服务端生成相应的JS代码，并且把返回的Content-type设置为application/javascript即可。
       在生成这个这个对应的JS代码的时候，就比较有讲究了，一般是调用客户端网页已经存在的JS函数。
*/

func main() {
	r := gin.Default()
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, gin.H{"wechat": "flysnow_org"})
	})

	// ========================================================================
	// 防止 JSONP 劫持，增加 while(1)，进行死循环不断覆盖数据

	a := []string{"1", "2", "3"}
	r.GET("/secureJson", func(c *gin.Context) {
		c.SecureJSON(200, a)
	})

	r.Run(":8080")
}
