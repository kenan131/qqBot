package test

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"helloGo/dto"
	"helloGo/entity"
	"helloGo/log"
	"helloGo/service"
	"os"
	"strings"
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
	})

	t.Run("test query table data", func(t *testing.T) {
		// 测试查询sql
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			t.Error(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT `id`, `key`, `reply` FROM fix_reply")
		if err != nil {
			t.Error(err)
		}
		defer rows.Close()

		for rows.Next() {
			var reply entity.FixReply
			if err := rows.Scan(&reply.Id, &reply.Key, &reply.Reply); err != nil {
				t.Error(err)
			}
			fmt.Println(reply)
		}
	})

	t.Run("test operation table data", func(t *testing.T) {
		// 测试插入sql
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			t.Error(err)
		}
		defer db.Close()

		var reply = entity.FixReply{
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
			t.Error(err)
		}
		fmt.Println(id)
	})
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
	t.Run("import idiom", func(t *testing.T) {
		connStr := "root:123456@tcp(127.0.0.1:3306)/qqbot?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Error("Error connecting to the database: ", err.Error())
		}
		defer db.Close()
		// 打开成语文件
		file, err := os.Open("idiom.txt")
		if err != nil {
			t.Error(err)
		}
		defer file.Close()

		// 创建扫描器
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines) // 按行分割

		// 逐行读取并导入数据库
		for scanner.Scan() {
			text := scanner.Text()
			splits := strings.Split(text, "、")
			for _, idiom := range splits {
				//time.Sleep(700 * time.Millisecond)
				if len(idiom) != 12 {
					continue
				}
				rows, err := db.Query("SELECT idiom from idiom_library where idiom = ?", idiom)
				if err != nil {
					t.Error(err)
				}
				if rows.Next() {
					var temp string
					if err := rows.Scan(&temp); err != nil {
						t.Error(err)
						return
					}
					if strings.TrimSpace(temp) != "" {
						// 数据存在，则不插入
						continue
					}
				}
				// 插入成语到数据库
				result, err := db.Exec("INSERT INTO idiom_library (`idiom`) VALUES (?)", idiom)
				if err != nil {
					t.Error(err)
				}
				// 获取插入操作的自增ID
				id, err := result.LastInsertId()
				if err != nil {
					t.Error(err)
				} else {
					fmt.Printf("成语 '%s' 已插入，ID: %d\n", idiom, id)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			t.Error(err)
		}
	})
}
