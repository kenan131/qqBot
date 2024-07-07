package service

import (
	"database/sql"
	"helloGo/dto"
	"helloGo/entity"
	"helloGo/log"
)

// GetFixReplyList 获取所有默认回复内容
func GetFixReplyList(db *sql.DB) ([]entity.FixReply, error) {
	var fixReplyList []entity.FixReply
	rows, err := db.Query("SELECT `id`, `key`, `reply` FROM fix_reply")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var reply entity.FixReply
		if err := rows.Scan(&reply.Id, &reply.Key, &reply.Reply); err != nil {
			log.Error(err)
		}
		fixReplyList = append(fixReplyList, reply)
	}
	return fixReplyList, nil
}

// GetFixReplyByKey 根据key获取默认回复内容
func GetFixReplyByKey(key string, db *sql.DB) (*entity.FixReply, error) {
	rows, err := db.Query("SELECT `id`, `key`, `reply` FROM fix_reply where `key` = ? ", key)
	if err != nil {
		return nil, err
	}
	var reply = &entity.FixReply{}
	for rows.Next() {
		if err2 := rows.Scan(&reply.Id, &reply.Key, &reply.Reply); err2 != nil {
			return nil, err2
		}
		break
	}
	if reply.Id == 0 {
		reply = nil
	}
	return reply, nil
}

// InsertFixReply 插入默认回复
func InsertFixReply(key string, reply string, db *sql.DB) error {
	temp, _ := GetFixReplyByKey(key, db)
	if temp != nil {
		// 不等于空则修改 使用id做where条件
		_, err := db.Exec("UPDATE fix_reply SET `reply` = ? WHERE `id` = ?", reply, temp.Id)
		if err != nil {
			return err
		}
	} else {
		_, err := db.Exec("INSERT INTO fix_reply (`key`, `reply`) VALUES (?, ?)", key, reply)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteFixReply 根据key删除默认回复
func DeleteFixReply(key string, db *sql.DB) (string, error) {
	temp, _ := GetFixReplyByKey(key, db)
	if temp != nil {
		// 不等于空则删除 使用id做where条件
		_, err := db.Exec("DELETE FROM fix_reply WHERE `id` = ?", temp.Id)
		if err != nil {
			return "删除失败！", err
		}
	} else {
		return "删除的key不存在哦！", nil
	}
	return "删除成功！", nil
}

// InitDefaultMap 初始化map
func InitDefaultMap(db *sql.DB) error {
	fixReplyList, err := GetFixReplyList(db)
	if err != nil {
		return err
	}
	for _, fixReply := range fixReplyList {
		dto.DefaultMessageMap[fixReply.Key] = fixReply.Reply
	}
	return nil
}
