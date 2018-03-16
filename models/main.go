package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	dbname = "ginrestdemo.db"
)

var (
	db     *gorm.DB
	tables []interface{}
)

func init() {
	var err error

	db, err = gorm.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}
	//db.SetMaxIdleConns(10)
	//db.SetMaxOpenConns(10)

	tables = append(tables, new(User))
	db.AutoMigrate(tables...)
}
