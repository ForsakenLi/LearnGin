package TemplateNest

import (
	"fmt"
	"html/template"
	"net/http"
	"testing"
)

func f1(w http.ResponseWriter, r *http.Request) {
	f2 := func(name string) (string, error) {
		return name + " adj!", nil
	}
	//定义模版
	t := template.New("f.tmpl")
	//告诉模板引擎，我现在多了一个自定义的函数f
	t.Funcs(template.FuncMap{
		"myFunc": f2,
	})
	//解析模版
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v\n", err)
		return
	}
	name := "abc"
	t.Execute(w, name)

	//渲染模版
}

func TestTemplate(t *testing.T) {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed, err: %v", err)
		return
	}
}
