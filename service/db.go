package service

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"runtime"
)

func GetDbConnectUrl(name string) (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		file := fmt.Sprintf("%s/%s", path.Dir(filename), name)
		var conf struct {
			ConnectUrl string `yaml:"connectUrl"`
		}
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		if err = yaml.Unmarshal(content, &conf); err != nil {
			return "", err
		}
		return conf.ConnectUrl, err
	}
	return "", nil
}

func GetDb(name string) (*sql.DB, error) {
	connectUrl, err := GetDbConnectUrl(name)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mysql", connectUrl)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// 测试连接
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
