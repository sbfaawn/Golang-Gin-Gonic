package configuration

import (
	"Golang-Gin-Gonic/model"
	"log"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	dsn := "root:root@tcp(localhost:3306)/gin-gonic?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.New(
		mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatalln("? Could not load environment variables", err)
	}

	err = DB.AutoMigrate(&model.Book{})
	fmt.Println("Table Migration is done")

	if err != nil{
		fmt.Println(err.Error())
	}
}
