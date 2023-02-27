package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseResponse struct {
	StatusCode int8   `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

var db *gorm.DB = InitDataSource()

func InitDataSource() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "project:project@tcp(rm-2ze62585lf96k7285mo.mysql.rds.aliyuncs.com:3306)/sqlxy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
	return db
}
