package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"test-go-worker/config"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	conn *gorm.DB
	err  error
)

func ConnectionMysql() *gorm.DB {
	config := config.NewViperConfig()

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.dbname"),
	)

	conn, err = gorm.Open("mysql", dbURI)
	if err != nil {
		fmt.Println("MYSQL DATABASE Cant Connect", err)
		panic(err)
	}

	return conn

}
