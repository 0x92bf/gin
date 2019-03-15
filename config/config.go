package config

import (
	"github.com/Unknwon/goconfig"
	"log"
)

type DbCfg struct {
	Username string
	Password string
	Database string
	Host string
	Port string
}

var Cfg *goconfig.ConfigFile

func init()  {
	var err error
	Cfg,err = goconfig.LoadConfigFile("./config/db.ini")
	if err != nil{
		log.Fatal(err)
	}
	return
}

func GetMysqlCfg() (mysql DbCfg) {
	mysql.Username,_ = Cfg.GetValue("mysql","username")
	mysql.Password,_ = Cfg.GetValue("mysql","password")
	mysql.Database,_ = Cfg.GetValue("mysql","database")
	mysql.Host,_ = Cfg.GetValue("mysql","host")
	mysql.Port,_ = Cfg.GetValue("mysql","port")
	return mysql
}