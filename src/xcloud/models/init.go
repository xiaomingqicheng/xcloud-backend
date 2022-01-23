package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/beego/beego/v2/client/orm"
	"github.com/astaxie/beego"
)

func init() {
	var mysql_connection string
	mysql_connection = beego.AppConfig.String("mysql")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:Root@123@tcp(192.168.0.108:3306)/xcloud?charset=utf8")
	orm.RegisterDataBase("default", "mysql", mysql_connection)
	orm.RegisterModel(new(Env),new(Cluster),new(Cert),new(Registry))
	orm.RunSyncdb("default", false, true)
}
