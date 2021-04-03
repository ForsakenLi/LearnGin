package GinStart

import (
	"fmt"
	"net/http"
	"testing"
	"html/template"
)

type User struct {
	Name string
	Gender string
	Age int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err: %v", err)
		return
	}
	//渲染模板
	u1 := User{
		Name:   "abc",
		Gender: "male",
		Age:    18,
	}
	err = t.Execute(w, u1)
	if err != nil {
		fmt.Printf("render template failed, err: %v", err)
		return
	}
}

func TestTemplate(t *testing.T) {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err: %v", err)
		return
	}
}
