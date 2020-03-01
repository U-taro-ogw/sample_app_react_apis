package db

import (
	"github.com/U-taro-ogw/gin_todo/src/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func Init() *gorm.DB {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_ROOT_PASSWORD")
	PROTOCOL := "tcp(db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		panic("データベース接続に失敗しました。")
	}
	db.AutoMigrate(&models.Todo{})

	return db
}
