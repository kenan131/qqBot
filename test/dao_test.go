package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"helloGo/dto"
	"helloGo/log"
	"helloGo/service"
	"testing"
)

func TestDao(t *testing.T) {
	t.Run("test fixReply Sql", func(t *testing.T) {
		// 测试 fixReply查询list方法
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()
		rows, err := service.GetFixReplyList(db)
		if err != nil {
			t.Error(err)
		}
		for _, value := range rows {
			fmt.Println(value)
		}
	})
	t.Run("test fixReply sql by Key", func(t *testing.T) {
		// 测试 fixReply 查询指定key方法
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()
		row, err := service.GetFixReplyByKey("有人吗？", db)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(row)
	})
	t.Run("test fixReply insert sql", func(t *testing.T) {
		// 测试从mqp中获取数据并插入到数据库中
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()
		for key, value := range dto.DefaultMessageMap {
			err1 := service.InsertFixReply(key, value, db)
			if err1 != nil {
				t.Error(err1)
			}
		}
	})
	t.Run("test initFixReply method", func(t *testing.T) {
		// 测试 初始化方法 从数据库查询数据存到map中。
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()
		service.InitDefaultMap(db)
		for key, value := range dto.DefaultMessageMap {
			fmt.Println(key + " " + value)
		}
	})
	t.Run("test delete fixReply method", func(t *testing.T) {
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()
		msg, err := service.DeleteFixReply("伤心", db)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(msg)
	})
	t.Run("test initIiom", func(t *testing.T) {
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()
		err1 := service.InitIdiomLibrary(db)
		if err1 != nil {
			t.Error(err1)
		}
		// 随机返回一个以"望"开头的成语
		result, ok := dto.StartsWithRandom("望")
		if ok {
			fmt.Printf("Random idiom starting with '望': %s\n", result)
		} else {
			fmt.Println("No idioms found starting with '望'")
		}
		result1, ok := dto.StartsWithRandom("望")
		if ok {
			fmt.Printf("Random idiom starting with '望': %s\n", result1)
		} else {
			fmt.Println("No idioms found starting with '望'")
		}
		result2, ok := dto.StartsWithRandom("望")
		if ok {
			fmt.Printf("Random idiom starting with '望': %s\n", result2)
		} else {
			fmt.Println("No idioms found starting with '望'")
		}
	})
}
