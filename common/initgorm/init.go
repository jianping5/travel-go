package initgorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitGorm gorm 初始化
func InitGorm(MysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{})
	if err != nil {
		panic("连接MySQL数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接MySQL数据库成功")
	}
	return db
}
