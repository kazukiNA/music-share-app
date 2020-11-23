package router

import (
	"music_app_api/controller"
	"music_app_api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupRouter() *gin.Engine {
	// jwt handler
	r := gin.Default()
	r.Use(middleware.SetHeaderHandler())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))
	r.Use(gin.Recovery())
	r.GET("/ping", controller.Ping)
	r.POST(("users"), controller.AddUsers)
	return r
}
