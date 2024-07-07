package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"
	"helloGo/dto"
	"helloGo/net"
	"helloGo/service"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

var process service.Processor

func main() {
	token, err := dto.GetToken(dto.ConfigStr)
	checkError(err)
	httpClient := net.MyHttpClient{}
	httpClient.InitHttpClient(token)
	// 获取websocket连接ip地址。
	webSocketIp := httpClient.GetMethod(dto.GetWebSocketIp)
	// 获取数据库连接
	db := GetDbConnect(dto.ConfigStr)
	// 初始化map数据。
	initData(db)
	process = service.Processor{
		Api: httpClient,
		Db:  db,
	}
	// 订阅事件
	intent := dto.IntentGuildAtMessage | dto.IntentGuildMessages | dto.IntentDirectMessages
	// 注册艾特消息 处理器
	net.RegisterHandler(dto.WSDispatchEvent, dto.EventAtMessageCreate, AtMessageHandler())
	// 注册普通消息 处理器
	net.RegisterHandler(dto.WSDispatchEvent, dto.EventMessageCreate, MessageHandler())
	// 注册私信 处理器
	net.RegisterHandler(dto.WSDispatchEvent, dto.EventDirectMessageCreate, DirectMessageHandler())
	container := net.New()
	// 启动webSocket
	container.Start(webSocketIp, token, intent)
}

// AtMessageHandler 艾特消息处理器
func AtMessageHandler() net.EventHandler {
	return func(event *dto.WSPayload, message []byte) error {
		data := &dto.Message{}
		if err := ParseData(message, data); err != nil {
			return err
		}
		if err := process.ProcessAtMessage(data); err != nil {
			return err
		}
		return nil
	}
}

// MessageHandler 普通消息处理器
func MessageHandler() net.EventHandler {
	return func(event *dto.WSPayload, message []byte) error {
		data := &dto.Message{}
		if err := ParseData(message, data); err != nil {
			return err
		}
		process.ProcessMessage(data)
		return nil
	}
}

func DirectMessageHandler() net.EventHandler {
	return func(event *dto.WSPayload, message []byte) error {
		data := &dto.Message{}
		if err := ParseData(message, data); err != nil {
			return err
		}
		//process.ProcessDirectMessage(data)
		return nil
	}
}

func ParseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")
	return json.Unmarshal([]byte(data.String()), target)
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

func initData(db *sql.DB) {
	// 使用sql 初始化数据
	err := service.InitDefaultMap(db)
	checkError(err)
	err = service.InitIdiomLibrary(db)
	checkError(err)
}
