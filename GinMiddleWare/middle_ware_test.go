package GinMiddleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func indexHandler(context *gin.Context) {
	name, ok := context.Get("name")
	if !ok {
		name = "Anonymous"
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

//定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	start := time.Now()
	c.Set("name", "abc")
	go func(c *gin.Context) {
		//在协程中只能使用context的拷贝！！
		//goroutine只能读取context的值，不能修改
	}(c.Copy())
	c.Next() //调用后续的处理函数
	//c.Abort()	//阻止调用后序的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Next() //调用后续的处理函数
	fmt.Println("m2 out...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或一些其他的准备工作
	return func(context *gin.Context) {
		if doCheck {
			//存放具体逻辑
			//是否登陆判断
			//if 是登陆用户
			//c.Next()
			//else
			//c.Abort()
		} else {
			context.Next()
		}
	}
}

func TestMidWare(t *testing.T) {
	r := gin.Default()

	r.Use(m1, m2) //全局注册中间件函数m1
	r.GET("/index", indexHandler)

	r.Run()
}
