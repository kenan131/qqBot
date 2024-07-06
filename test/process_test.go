package test

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v3"
	"helloGo/dto"
	"helloGo/net"
	"helloGo/service"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	t.Run("test split", func(t *testing.T) {
		// 测试分割字符串
		s := "valueBef:oreColon:valueAfterColon:1121121:98989"
		// 根据字符串的第一个分隔符进行切割
		firstColonParts := strings.SplitN(s, ":", 2)
		fmt.Println("第一个冒号前的值:", firstColonParts[0])
		fmt.Println("第一个冒号之后的所有值:", firstColonParts[1])
	})
	t.Run("test Instructions003", func(t *testing.T) {
		token, err := dto.GetToken("config.yaml")
		if err != nil {
			t.Error(err)
		}
		httpClient := net.MyHttpClient{}
		httpClient.InitHttpClient(token)
		// 获取数据库连接
		db := GetDbConnect("config.yaml")
		// 初始化map数据。
		err1 := service.InitDefaultMap(db)
		checkError(err1)
		process := service.Processor{
			Api: httpClient,
			Db:  db,
		}
		res := process.Instructions003("指令:回复<emoji:33>")
		if res != "添加成功，赶快来试试吧！" {
			t.Error()
		}
	})

	t.Run("test Instructions004 ", func(t *testing.T) {
		token, err := dto.GetToken("config.yaml")
		if err != nil {
			t.Error(err)
		}
		httpClient := net.MyHttpClient{}
		httpClient.InitHttpClient(token)
		// 获取数据库连接
		db := GetDbConnect("config.yaml")
		// 初始化map数据。
		err1 := service.InitDefaultMap(db)
		checkError(err1)
		process := service.Processor{
			Api: httpClient,
			Db:  db,
		}
		res := process.Instructions004("指令")
		if res != "删除成功！" {
			t.Error()
		}
	})
}
func GetDbConnect(name string) *sql.DB {
	connectUrl, err := GetDbConnectUrl(name)
	if err != nil {
		log.Fatalln(err)
	}
	db, err := service.GetDb(connectUrl)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

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

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
