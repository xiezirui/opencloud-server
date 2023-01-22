package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"opencloud-server/model/dao"
)

// init DB connect to mysql
func InitDatabase() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/opencloud_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("========>链接失败")
	}

	return db
}

// finally close connection
func CloseDatabase(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

// sync struct to table
func SyncStruct(db *gorm.DB) {
	err := db.AutoMigrate(&dao.User{})
	err = db.AutoMigrate(&dao.Blog{})
	if err != nil {
		fmt.Println("error")
	}
}
