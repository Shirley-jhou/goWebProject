package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginController struct {
	SqlDB *gorm.DB
}

//登入判斷
func (*LoginController) LoginConfirm(c *gin.Context)  {
	//TODO:
	c.Writer.WriteString("LoginConfirm")
	fmt.Println("-----------------login----------------")
}