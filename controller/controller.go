package controller

import (
	"music_app_api/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Ping(c *gin.Context) {
	c.Status(200)
}

func DB(c *gin.Context) *gorm.DB {
	return c.MustGet("db").(*gorm.DB)
}

func CurrentUser(c *gin.Context) *model.User {
	return c.MustGet("currentUser").(*model.User)
}
