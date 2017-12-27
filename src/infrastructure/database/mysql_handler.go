package database

import (
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"strings"
	"fmt"
	"log"
	"os"
)

type MysqlHandler struct {
	Connection *gorm.DB
}

func NewMysqlHandler() *MysqlHandler {

	db, err := gorm.Open("mysql", makeDSN())

	if err != nil {
		log.Fatalln("Failed to connect database.")
	}

	if os.Getenv("QUERY_LOG") == "true" {
		db.LogMode(true)
		db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}

	handler := new(MysqlHandler)
	handler.Connection = db

	return handler
}

func GetFullDSN() string {
	return "mysql://" + makeDSN()
}

func makeDSN() string {

	host     := os.Getenv("MYSQL_HOST")
	port     := os.Getenv("MYSQL_PORT")
	user     := os.Getenv("MYSQL_USER")
	pass     := os.Getenv("MYSQL_PASS")
	dbName   := os.Getenv("MYSQL_DB_NAME")
	protocol := os.Getenv("MYSQL_PROTOCOL")
	dbArgs   := os.Getenv("MYSQL_ARGS")

	if strings.Trim(dbArgs, " ") != "" {
		dbArgs = "?" + dbArgs
	} else {
		dbArgs = ""
	}

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, pass, protocol, host, port, dbName, dbArgs)
}

func (handler *MysqlHandler) FindByID(result interface{}, id int) (error) {
	return handler.Connection.Where("id = ?", id).Find(result).Error
}

func (handler *MysqlHandler) Close() {
	handler.Connection.Close()
}
