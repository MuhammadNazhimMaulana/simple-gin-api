package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMysql() *gorm.DB {

	dsn := "root:@tcp(localhost:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// If Error Exist
	if err != nil {
		panic(err.Error())
	}

	return db
}
