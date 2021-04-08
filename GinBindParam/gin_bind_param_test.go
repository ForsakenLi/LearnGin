package GinBindParam

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func TestBind(t *testing.T) {
	r := gin.Default()
	r.GET("/user", func(context *gin.Context) {
		//username := context.Query("username")
		//password := context.Query("password")
		//u := UserInfo{
		//	username: username,
		//	password: password,
		//}
		var u UserInfo
		//ShouldBind应传递指针形式的结构体
		//如果querystring中有对应参数就赋给该结构体对应名称的值
		//ShouldBind内部会通过反射来获取结构体对应的参数
		//因此结构体中的变量首字母需要大写
		//如果前端要求的首字母为小写的话
		//则需要通过`form:"xxx"`来指定链接里的名称
		//testURL:http://localhost:9090/user?username=abc&password=123456
		err := context.ShouldBind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.POST("/form", func(context *gin.Context) {
		var u UserInfo
		err := context.ShouldBind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			context.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":9090")
}
