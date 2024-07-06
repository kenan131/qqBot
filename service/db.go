package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDb(connectUrl string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectUrl)
	if err != nil {
		return nil, err
	}
	// 测试连接
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
