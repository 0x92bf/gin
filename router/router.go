package router

import (
	"github.com/gin-gonic/gin"
	"sas/common"
	"sas/controller"
)

func Register(router *gin.Engine) {
	mobile := router.Group("mobile", getHeaders)
	{
		mobile.POST("/merchant/register", controller.MctRegister)
	}
}

func getHeaders(c *gin.Context) {
	openid := c.GetHeader("openid")
	if openid == "" {
		common.ResponJson(c, "缺少openid", "", 0)
		return
	}
	c.Set("openid", openid)
	c.Next()
}
