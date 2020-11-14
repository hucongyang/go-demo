package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"

	"github.com/hucongyang/go-demo/conf"
)

// 用于models的初始化使用

var db *gorm.DB
var err error

type GORMBase struct {
	ID         int   `json:"id" gorm:"primary_key"`
	CreatedOn  int64 `json:"created_on"`
	ModifiedOn int64 `json:"modified_on"`
}

func init() {
	database := conf.Config().Database
	db, err = gorm.Open(database.Driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		database.User,
		database.Password,
		database.Host,
		database.Name))
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return database.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
