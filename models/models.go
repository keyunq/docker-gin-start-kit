package models

import (
	"fmt"
	"log"
	"userInfoService/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
		port                                 int
	)

	dbType = setting.Cfg.DBType
	dbName = setting.Cfg.Mysql.DbName
	user = setting.Cfg.Mysql.UserName
	password = setting.Cfg.Mysql.Password
	host = setting.Cfg.Mysql.Ip
	port = setting.Cfg.Mysql.Port
	debug := setting.Cfg.Mysql.Debug

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	fmt.Printf("LogMode : %+v", debug)
	db.SingularTable(true)
	db.LogMode(debug)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
