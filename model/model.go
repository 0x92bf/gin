package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sas/config"
)

var Db *gorm.DB
func init(){
	var err error
	mysql := config.GetMysqlCfg()
	jdbc := mysql.Username+":"+mysql.Password+"@tcp("+mysql.Host+":"+mysql.Port+")/"+mysql.Database+"?charset=utf8&parseTime=True&loc=Local"
	Db,err = gorm.Open("mysql",jdbc)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	if err != nil{
		log.Fatal(err)
	}
	return
}

