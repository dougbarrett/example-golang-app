package mysql

import (
	"fmt"

	"github.com/dougbarrett/example-golang-app"
	"github.com/dougbarrett/example-golang-app/storage"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysql struct {
	db *gorm.DB
}

func NewFrontend(
	username string,
	password string,
	host string,
	database string,
) (
	storage.Service,
	error,
) {
	m := &mysql{}
	connect := fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true", username, password, host, database)

	db, err := gorm.Open("mysql", connect)

	if err != nil {
		return nil, err
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&guestbook.Entry{}).Error

	if err != nil {
		return nil, err
	}

	m.db = db
	return m, nil
}
