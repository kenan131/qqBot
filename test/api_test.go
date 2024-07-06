package test

import (
	"fmt"
	"helloGo/dto"
	"helloGo/net"
	"log"
	"testing"
)

func TestGetReq(t *testing.T) {
	t.Run(
		"get websocket accessIp by gateway", func(t *testing.T) {
			token, err := dto.GetToken("config.yaml")
			// 从配置文件中获取token信息
			if err != nil {
				t.Error(err)
			}
			httpClient := net.MyHttpClient{}
			httpClient.InitHttpClient(token)
			// 获取websocket连接ip地址。
			ip := httpClient.GetMethod(dto.GetWebSocketIp)
			fmt.Println("从网关中获取的websocket连接ip：" + ip)
		},
	)
	t.Run("post method", func(t *testing.T) {
		token, err := dto.GetToken("config.yaml")
		// 从配置文件中获取token信息
		if err != nil {
			log.Fatalln(err)
		}
		httpClient := net.MyHttpClient{}
		httpClient.InitHttpClient(token)
		replyMessage := &dto.MessageToCreate{
			Content: "测试发送消息！",
		}
		httpClient.PostMethodParam(dto.MessagesURI, "channel_id", "654900577", replyMessage)
	})
}
