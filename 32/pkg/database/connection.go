package database

import (
	"fmt"
	"time"

	"skillbox/pkg/logs"
	"skillbox/pkg/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func init() {
	logs.InitialSet("main")
}

func Initial() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=users port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{}, &model.Friends{})
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Initial()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("ДБ не отвечает. Подождите %d секунд.", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Initial()
		}
	}
	return dbase
}
