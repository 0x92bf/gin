package common

import (
	"github.com/gin-gonic/gin"
)

func ResponJson(c *gin.Context,msg string,data interface{},errorCode interface{})  {
	c.JSON(200,gin.H{"msg":msg,"data":data,"errorCode":errorCode})
}
