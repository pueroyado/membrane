package server

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func NewMysqlConnect() *sqlx.DB {
	log.Println("start connect mysql")

	str := "%s:%s@tcp(%s)/%s"
	connectionString := fmt.Sprintf(str,
		"mysql",
		"mysql",
		"app-mysql:3306",
		"demo",
	)

	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		log.Fatal("Mysql connect error ", err)
	}
	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
