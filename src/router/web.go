package router

import (
	"Test/src/controllers"
	"Test/src/middlewares"
	"github.com/gin-gonic/gin"
)

func InitWebRouter (engine *gin.Engine){
	var (
		LoginController *controllers.LoginController
	)

	//接口訪問
	webapi := engine.Group("/test/v1" , middlewares.Ipblack)
	{
		webapi.POST("/Login" , LoginController.LoginConfirm)
	}
}