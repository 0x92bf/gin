package controller

import (
	_"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sas/model"
	"sas/common"
)


func MctRegister(c *gin.Context)  {
	var json model.MerchantRegisterInfo
	if err := c.ShouldBindJSON(&json);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err,merchant := model.CreateMerchant(c,json)
	if err !=nil{
		common.ResponJson(c,err.Error(),"",-1)
		return
	}
	common.ResponJson(c,"商户创建成功",merchant,0)
}
