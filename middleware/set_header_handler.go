package middleware

import (
	"fmt"
	"music_app_api/config"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SetHeaderHandler is set db session to context
func SetHeaderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		conf := config.Configuration{
			"mysql", "root", "hogehoge", "tcp(mysql:3306)", "music_app",
		}
		dbConnection := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", conf.USER, conf.PASS, conf.PROTOCOL, conf.DBNAME)
		db, err := gorm.Open(conf.DBMS, dbConnection)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
	}
}
