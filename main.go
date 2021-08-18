package main

import (
	"WebApi/Controller"
	"WebApi/FrameWork"
	"fmt"
	"net/http"
	"time"
)

func main()  {
	server := &http.Server{
		Addr:        ":8081",
		Handler:     FrameWork.Router,
		ReadTimeout: 5 * time.Second,
	}
	RegiterRouter(FrameWork.Router)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("start server error")
	}
	fmt.Println("start server success")
}

func RegiterRouter(handler *FrameWork.RouterHandler) {
	new(Controller.TestConterller).Router(handler)
}