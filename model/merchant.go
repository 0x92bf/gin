package model

import (
	//"errors"
	//"log"
	"time"
	"github.com/gin-gonic/gin"
)

type Merchant struct {
	Merchant_id int32 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Openid string
	Merchant_name string
	Linkman string
	Linkmobile string
	License_image string
	Tmt_name string `gorm:"-"`
	Create_time string
	Update_time string
	Is_del int32
}
type MerchantRegisterInfo struct {
	Merchant_name string
	Linkman string
	Linkmobile string
}

func (Merchant) TableName() string  {
	return "merchant"
}

func CreateMerchant(c *gin.Context,json MerchantRegisterInfo)(err error,merchant Merchant){
	//查询商户名是否存在
	//var hasMerchant Merchant
	//Db.Where("merchant_name = ?",json.Merchant_name).First(&hasMerchant)
	//if hasMerchant.Merchant_id != 0{
	//	return errors.New("商户名已存在"),hasMerchant
	//}
	openid := c.MustGet("openid").(string)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	merchant = Merchant{Openid:openid,Merchant_name:json.Merchant_name,Linkman:json.Linkman,Linkmobile:json.Linkmobile,Create_time:timeStr}
	result := Db.Omit("Merchant_id","License_image","Tmt_name","Update_time","Is_del").Create(&merchant)
	err = result.Error
	return
}