package goninja

import (
	 "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Model struct {
	Id int
	Created_at string
	Updated_at string
}


var DB, CONNECTION_ERROR = gorm.Open("postgres", "user=alex dbname=goblog_dev sslmode=disable")

func InitDB() {

	if CONNECTION_ERROR == nil {
		DB.DB()
		DB.DB().Ping()
		DB.DB().SetMaxIdleConns(10)
		DB.DB().SetMaxOpenConns(100)
	}

}



