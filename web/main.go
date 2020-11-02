package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", root)
	r.GET("/test", test)
	r.Run(":8080")
}

func root(c *gin.Context) {
	c.JSON(200, gin.H{
		"Blog":   "www.flysnow.org",
		"wechat": "flysnow_org",
	})
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"Blog":   "test",
		"wechat": "test",
	})
}
