package Controller

import (
	. "WebApi/FrameWork"
	. "WebApi/Model"
	. "WebApi/Services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestConterller struct {
}

//POST Content-Type=application/x-www-form-urlencoded
func (p *TestConterller) Insert(writer http.ResponseWriter, request *http.Request) {
	test := new(Test)
	fmt.Println("Name", request.PostFormValue("Name"))
	test.Name = request.PostFormValue("Name")
	json := InsertTest(*test)
	ResultJsonOk(writer, json)
	return
}

func (p *TestConterller) Insert1(c *gin.Context) {
	test := new(Test)
	fmt.Println("Name", c.PostForm("Name"))
	test.Name = c.PostForm("Name")
	c.JSON(200, InsertTest(*test))
}

func (p *TestConterller) Router(router *RouterHandler) {
	router.Router("/Insert", p.Insert)
}
