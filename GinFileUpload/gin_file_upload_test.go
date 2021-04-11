package GinFileUpload

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestUpload(t *testing.T) {
	r := gin.Default()
	r.LoadHTMLFiles("./")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)

	})
	r.Run(":9090")
}
