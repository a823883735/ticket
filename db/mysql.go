package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"ticket/conf"
)

var mysqlEngine *xorm.Engine

func MysqlInit() {
	mysqlSection := conf.Conf().Section("mysql")
	if !mysqlSection.HasKey("host") {
		log.Fatal("mysql load failed from the file init.go.")
	}
	var err error
	mysqlEngine, err = xorm.NewEngine(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			mysqlSection.Key("name"),
			mysqlSection.Key("pwd"),
			mysqlSection.Key("host"),
			mysqlSection.Key("port"),
			mysqlSection.Key("db"),
		),
	)
	if err != nil {
		log.Fatal("mysql connection failed. err: ", err)
	}
	mysqlEngine.ShowSQL(true)
	// 设置最大连接数
	mysqlEngine.SetMaxIdleConns(200)
}

// MYSQL returns this mysql db operation engine.
func MYSQL() *xorm.Engine {
	return mysqlEngine
}