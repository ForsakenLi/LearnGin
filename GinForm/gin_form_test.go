package GinForm

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestForm(t *testing.T) {
	r := gin.Default()
	//注意这里login和登陆完成的html都需要加载！
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})
	//处理客户端发来的post请求
	r.POST("/login", func(context *gin.Context) {
		username := context.DefaultPostForm("username","somebody")
		password := context.DefaultPostForm("password", "xxxx")
		//也可以使用GetPostForm形式，用ok来判断是否返回成功
		context.HTML(http.StatusOK, "index.html", gin.H{
			"Name": username,
			"Password": password,
		})
	})
	r.Run(":9090")
}
