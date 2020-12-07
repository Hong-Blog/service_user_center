package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var Db *sqlx.DB

func init() {
	dataSource := os.Getenv("DATA_SOURCE")
	if dataSource == "" {
		dataSource = "root:123qwe@tcp(127.0.0.1:3306)/hongblog?parseTime=true"
	}
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		log.Panicln("db err: ", err.Error())
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	Db = db
}
