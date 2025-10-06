package main

func MySQLCfg() {

	if Init_MySQL {

		FolderCheck("dao/mysql", "dao/mysql", "[MYSQL] ")
		WriteContentToConfigYaml(MySQL_Init_Content, "dao/mysql/mysql.go", "[MYSQL] ")
		WriteContentToConfigYaml(MySQL_Config_Yaml, "config.yaml", "[MYSQL] ")
	}
}

var (
	MySQL_Config_Yaml = `mysql: 
  host: "127.0.0.1"
  port: 3306
  user: "root"
  password: "root"
  dbname: "sql_demo"
  max_open_conns: 20
  max_idle_conns: 10`
	MySQL_Init_Content = `package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err: %v\n ", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return

}

func Close() {
	err := db.Close()
	if err != nil {
		zap.L().Error("Close Mysql DB failed, err: %v\n ", zap.Error(err))
	}
}
`
)
