package cmd

import (
	"Test/src/config"
	"Test/src/database"
	"Test/src/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init()  {

}

func Execute()  {
	//設置時區UTC+8時間
	location , err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("設置時區錯誤" + err.Error())
		return
	}

	time.Local = location

	//初始化config
	config.Setup()

	//初始化db
	database.Setup()

	//初始化gin
	engine := gin.New()
	engine.Use(gin.Logger(),gin.Recovery())

	//配置接口
	router.InitWebRouter(engine)

	baseServer := "0.0.0.0:" + fmt.Sprintf("%d", config.Configs.TestAdmin.Port)
	srv := &http.Server{
		Addr:    baseServer,
		Handler: engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("服务监听: " + err.Error())
		}
	}()

}
