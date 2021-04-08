package GinJson

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestJson(t *testing.T) {
	r := gin.Default()
	r.GET("/json", func(context *gin.Context) {
		//way1 : using map
		data := map[string]interface{}{
			"name" : "abc",
			"message" : "hello world!",
			"age" : 18,
		}
		context.JSON(http.StatusOK, data)
	})

	r.GET("/json2", func(context *gin.Context) {
		//way2 : using gin.H (equal to map[string]interface{})
		data2 := gin.H{
			"name" : "abc",
			"message" : "hello world!",
			"age" : 18,
		}
		context.JSON(http.StatusOK, data2)
	})
	r.GET("/json3", func(context *gin.Context) {
		//way3 : using structure
		type msg struct {
			//结构体中内容必须首字母大写，否则外部无法读取
			//可以写tag来指定传递给前端时使用的名称
			Name string `json:"name"`
			Message string	`json:"message"`
			Age int	`json:"age"`
		}
		context.JSON(http.StatusOK, msg{
			Name:    "abc",
			Message: "hello world!",
			Age:     18,
		})
	})
	r.Run(":9090")
}
