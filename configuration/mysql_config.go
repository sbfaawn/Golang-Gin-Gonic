package configuration

import (
	gormLogger "Golang-Gin-Gonic/logger"
	"Golang-Gin-Gonic/model"
	"fmt"
	"log"

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
		&gorm.Config{
			Logger: gormLogger.NewGormLogger(),
		},
	)

	if err != nil {
		log.Fatalln("? Could not load environment variables", err)
	}

	err = DB.AutoMigrate(&model.Book{})
	fmt.Println("Error DB Migration : ", err)
	fmt.Println("Table Migration is done")

	isEmpty := false

	if isEmpty {
		populateData(DB)
	}

	if err != nil {
		fmt.Println(err.Error())
	}
}

func populateData(db *gorm.DB) {
	db.CreateInBatches([]model.Book{
		{
			Title:  "How to be Millionaire",
			Author: "Good Man Stainley",
		},
		{
			Title:  "Big Guy, Big Dong",
			Author: "Grey Outmilk",
		},
		{
			Title:  "Maddeline, Where are you now?",
			Author: "Bang John Halal",
		},
	}, 3)

	fmt.Println("Data has been Populated!!!!")
}
