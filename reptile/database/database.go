package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var (
	DB *gorm.DB //DB对象
)

func init() {

	//初始化MYSQL
	initMYSQL(fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_USER_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	))
}

//初始化MySQL
func initMYSQL(dataSource string) {
	fmt.Println("Init MYSQL.......")

	var err error
	DB, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Info),
		CreateBatchSize: 100,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Init MYSQL Success.")
}
