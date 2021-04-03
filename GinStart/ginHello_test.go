package GinStart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestGin1(t *testing.T) {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/book", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
