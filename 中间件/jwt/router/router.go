package router

import (
	"github.com/gin-gonic/gin"
	"jwt/controller"
	"jwt/middelware"
)

func Router(r *gin.Engine) {

	//用户登录
	r.GET("/login", controller.LoginController)
	//使用中间件
	r.Use(middelware.JWTAuth())
	//获取列表数据
	r.GET("/list", controller.UserListController)
}
