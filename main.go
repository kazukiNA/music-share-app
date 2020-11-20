package main

import (
	"music_app_api/router"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")

}
