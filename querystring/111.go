package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Hobby  string
	Lover  string
}

var db2 *sqlx.DB
var logger *zap.Logger

func main() {
	viper.SetDefault("file", "./")
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&UserInfo{})
	fmt.Println()
	//u1 := UserInfo{2, "mg", "女", "eat", "phm"}
	//db.Create(&u1)
	var u UserInfo
	db.First(&u)
	db.First(new(UserInfo))
	db.Table("aaa").CreateTable()
	db.Model(&u).Update("Gender", "female")

	// r.GET("/web", func(context *gin.Context) {
	//	name := context.Query("que  ry")
	//	context.JSON(http.StatusOK, gin.H{
	//		"name": name,
	//	})
	//})
	//r.LoadHTMLFiles("./querystring/login.html")
	//r.GET("/login", func(c *gin.Context) {
	//	go func() {}()
	//	c.HTML(http.StatusOK,"login.html", nil)
	//})
	//
	//r.Run(":8080")
}
