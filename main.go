package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	DBMigrate(DBConnect())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	r.Run(":8080")
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	return db
}

func DBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "hogehoge"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "music_app"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}
