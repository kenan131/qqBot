package main

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"helloGo/dto"
	"helloGo/net"
	"helloGo/service"
	"log"
)

var configStr = "config.yaml"
var process service.Processor

func main() {
	token, err := dto.GetToken(configStr)
	if err != nil {
		log.Fatalln(err)
	}
	httpClient := net.MyHttpClient{}
	httpClient.InitHttpClient(token)
	// 获取websocket连接ip地址。
	webSocketIp := httpClient.GetMethod(dto.GetWebSocketIp)
	// 获取数据库连接
	db, err := service.GetDb(configStr)
	defer db.Close()
	process = service.Processor{
		Api: httpClient,
		Db:  db,
	}
	// 订阅事件
	intent := dto.IntentGuildAtMessage | dto.IntentGuildMessages
	// 注册艾特消息 处理器
	net.RegisterHandler(dto.WSDispatchEvent, dto.EventAtMessageCreate, AtMessageHandler())
	net.RegisterHandler(dto.WSDispatchEvent, dto.EventMessageCreate, MessageHandler())
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

func ParseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")
	return json.Unmarshal([]byte(data.String()), target)
}
