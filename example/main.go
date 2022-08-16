package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/go-woo/protoc-gen-gin/example/v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1.RegisterGreeterRouter(r)
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Server starting error:%v\n", err)
	}
}
