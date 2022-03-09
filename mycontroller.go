package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type MyController struct {
	Ctx iris.Context
}

//  会在控制器之前被调用
func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	// method
	// Path
	// The controller's BeforeActivation
	// hook is called before the controller is activated.
	b.Handle("GET", "/test/{id:long}", "GetHello")
}

func (m *MyController) GetHello() interface{} {
	return iris.Map{
		"message": "Hello",
	}
}

// 基于RESTfulAPI的路由
// method: get
// path: /root
func (m *MyController) Get() string {
	return "hey from my controller"
}

// method: get
// path: /root/id
func (m *MyController) GetBy(id int64) interface{} {
	return iris.Map{
		"id":      id,
		"message": "hey from my controller",
	}
}

// method: get
// path: /root/hello/word
func (m *MyController) GetHelloWorld() interface{} {
	return iris.Map{
		"message": "Hello World",
	}
}

// method: any
// path: /root/test
func (m *MyController) AnyTest() {
	m.Ctx.WriteString("test")
}
