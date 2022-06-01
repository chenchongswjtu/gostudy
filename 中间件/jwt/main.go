package main

import (
	"github.com/gin-gonic/gin"
	"jwt/router"
)

func main() {
	r := gin.Default()
	//路由
	router.Router(r)
	r.Run()
}
