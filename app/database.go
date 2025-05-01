package app

import (
	"bubblevy/restful-api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:CsnHHCgrFrnFFCtgEdVxBtYuzHwWHQYE@tcp(interchange.proxy.rlwy.net:50548)/railway")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
