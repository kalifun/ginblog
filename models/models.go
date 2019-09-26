package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"kalifun/ginblog/pkg/setting"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	getdb, err := setting.Conf.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	dbType = getdb.Key("TYPE").String()
	dbName = getdb.Key("NAME").String()
	user = getdb.Key("USER").String()
	password = getdb.Key("PASSWORD").String()
	host = getdb.Key("HOST").String()
	tablePrefix = getdb.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	//结构体为非复数
	db.SingularTable(true)
	//设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	//设置最大打开的连接数
	db.DB().SetMaxOpenConns(100)
}

func CloseDB()  {
	defer db.Close()
}