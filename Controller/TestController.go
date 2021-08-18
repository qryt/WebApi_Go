package Controller

import (
	. "WebApi/FrameWork"
	. "WebApi/Model"
	. "WebApi/Services"
	"fmt"
	"net/http"
)
type TestConterller struct {

}

//POST Content-Type=application/x-www-form-urlencoded
func (p *TestConterller)Insert (writer http.ResponseWriter,request *http.Request) {
	test:=new(Test)
	fmt.Println("Name",request.PostFormValue("Name"))
	test.Name=request.PostFormValue("Name")
	json:=InsertTest(*test)
	ResultJsonOk(writer,json)
	return
}

func (p *TestConterller) Router(router *RouterHandler) {
	router.Router("/Insert", p.Insert)
}



