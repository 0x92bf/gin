package main

import (
	"github.com/gin-gonic/gin"
	"github.com/braintree/manners"
	_"log"
	"sas/controller"
	"sas/common"
)

func main()  {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	mobile := router.Group("mobile",getHeaders)
	{
		mobile.POST("/merchant/register",controller.MctRegister)
	}

	manners.ListenAndServe(":81",router)
	manners.Close()
}

func getHeaders(c *gin.Context){
	openid := c.GetHeader("openid")
	if openid == ""{
		common.ResponJson(c,"缺少openid","",0)
		return
	}
	c.Set("openid",openid)
	c.Next()
}