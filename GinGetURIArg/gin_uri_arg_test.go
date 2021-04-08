package GinGetURIArg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestURIArg(t *testing.T) {
	r := gin.Default()

	//匹配的项目在同一个函数中应该一样长，如果出现/blog/:year/:month"和"/:name/:age"就会出现匹配无法判定的情况（会报错）
	r.GET("/blog/:year/:month", func(context *gin.Context) {
		year := context.Param("year")
		month := context.Param("month")
		context.JSON(http.StatusOK, gin.H{
			"year" : year,
			"month" : month,
		})
	})
	r.GET("/user/:name/:age", func(context *gin.Context) {
		//获取路径参数
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})
	})

	r.Run(":9090")
}
