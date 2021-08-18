package main

import (
	. "WebApi/Controller"
	"WebApi/FrameWork"
	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	control := new(TestConterller)
	c.GET("Insert1", control.Insert1)
	c.Run("localhost:8081")
}

func RegiterRouter(handler *FrameWork.RouterHandler) {
	new(TestConterller).Router(handler)
}
