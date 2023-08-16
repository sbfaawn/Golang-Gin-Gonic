package configuration

import (
	gormLogger "Golang-Gin-Gonic/logger"
	"Golang-Gin-Gonic/model"
	properties_reader "Golang-Gin-Gonic/properties/reader"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	conf := properties_reader.Config.Database.ConfigMySql
	var err error

	dsn := fmt.Sprint(conf.Username, ":", conf.Password, "@tcp(", conf.Address, ":", conf.Port, ")/", conf.Database, "?charset=utf8&parseTime=True&loc=Local")
	// dsn := "root:root@tcp(localhost:3306)/gin-gonic?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.New(
		mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			Logger: gormLogger.NewGormLogger(),
		},
	)

	if err != nil {
		log.Fatalln("? : Could Established Connection to Databases", err)
	}

	if conf.Migrate {
		err = DB.AutoMigrate(&model.Book{}, &model.Credential{})
		fmt.Println("Error DB Migration : ", err)
		fmt.Println("Table Migration is done")
	}

	if conf.Populated {
		populateData(DB)
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

	db.CreateInBatches([]model.Credential{
		{Username: "admin", Password: "Admin123"},
		{Username: "dhika", Password: "Dhika78ty"},
	}, 2)

	fmt.Println("Data has been Populated!!!!")
}
