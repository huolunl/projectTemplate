/*
   @Author:huolun
   @Date:2021/11/22
   @Description
*/
package genericAPIserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RouterRegister(g *gin.Engine)  {
	g.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
}
