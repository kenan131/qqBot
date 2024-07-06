package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"helloGo/entity"
	"helloGo/log"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("test connect mysql dataBase", func(t *testing.T) {
		// 测试数据库连接
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqBot?charset=utf8mb4&parseTime=True&loc=Local"

		// 打开数据库连接
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()

		// 测试连接
		err = db.Ping()
		if err != nil {
			log.Error("Error pinging the database: ", err.Error())
		}

		fmt.Println("Successfully connected to the database.")
	})

	t.Run("test query table data", func(t *testing.T) {
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"

		// 打开数据库连接
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()
		rows, err := db.Query("SELECT `id`, `key`, `reply` FROM fix_reply")
		if err != nil {
			log.Error(err)
		}
		defer rows.Close()

		for rows.Next() {
			var reply entity.DefaultReply
			if err := rows.Scan(&reply.Id, &reply.Key, &reply.Reply); err != nil {
				log.Error(err)
			}
			fmt.Println(reply)
		}
	})

	t.Run("test operation table data", func(t *testing.T) {
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"

		// 打开数据库连接
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()

		var reply = entity.DefaultReply{
			Key:   "有人吗？",
			Reply: "大哥，机器人小弟我在！<emoji:16>",
		}
		res, err := db.Exec("INSERT INTO fix_reply (`key`, `reply`) VALUES (?, ?)", reply.Key, reply.Reply)
		if err != nil {
			log.Error(err)
			return
		}

		// 获取新插入行的ID
		id, err := res.LastInsertId()
		if err != nil {
			log.Error(err)
		}
		fmt.Println(id)
	})
}
