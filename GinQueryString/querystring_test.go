package GinQueryString

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestQueryString(t *testing.T) {
	r := gin.Default()
	r.GET("/web", func(context *gin.Context) {
		//获取浏览器发请求携带的query string携带的参数
		//Query的参数表示请求中对应参数的key
		//http://localhost:9090/web?query=xxx
		name := context.Query("name")
		age := context.Query("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})
	})

	r.GET("/web2", func(context *gin.Context) {
		//如果key参数缺失，使用DefaultQuery可以补全缺省值
		name := context.DefaultQuery("query","somebody")
		context.JSON(http.StatusOK, name)
	})

	r.GET("/web3", func(context *gin.Context) {
		//way3
		name, ok := context.GetQuery("query")
		if !ok {
			//取不到
			name = "somebody"
		}
		context.JSON(http.StatusOK, name)
	})

	r.Run(":9090")
}
