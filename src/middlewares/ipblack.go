package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Ipblack(c *gin.Context){
	ip := c.ClientIP()
	fmt.Println("訪問ip:", ip)
}
