package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectGORM() *gorm.DB {
	DBMS := "postgres"
	HOST := "localhost"
	PORT := "5000"
	USER := "postgres"
	DBNAME := "postgres"
	PASS := "password"

	CONNECT := "host=" + HOST + " port=" + PORT + " user=" + USER + " dbname=" + DBNAME + " password=" + PASS
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
