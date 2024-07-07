package service

import (
	"database/sql"
	"helloGo/dto"
	"helloGo/entity"
	"helloGo/log"
)

func InitIdiomLibrary(db *sql.DB) error {
	rows, err := db.Query("SELECT `id`, `idiom` FROM idiom_library")
	if err != nil {
		log.Error("初始化成语库失败！" + err.Error())
		return err
	}
	for rows.Next() {
		var idiom entity.IdiomLibrary
		if err := rows.Scan(&idiom.Id, &idiom.Idiom); err != nil {
			log.Error(err)
			return err
		}
		dto.Insert(idiom.Idiom)
	}
	return nil
}
