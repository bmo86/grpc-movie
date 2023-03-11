package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Instance_conn struct {
	db *gorm.DB
}

func Conn(url string) (*Instance_conn, error) {
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		return nil, err
	}

	return &Instance_conn{db}, err
}
