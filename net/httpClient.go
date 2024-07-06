package net

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	token "helloGo/dto"
	"log"
	"time"
)

// SocketIpStruct 请求网关返回的响应体
type SocketIpStruct struct {
	URL string `json:"url"`
}

type MyHttpClient struct {
	client *resty.Client
}

var domain = "https://sandbox.api.sgroup.qq.com"

func (m *MyHttpClient) InitHttpClient(token *token.Token) {
	authToken := token.GetString()
	// 创建一个新的 resty 客户端
	m.client = resty.New().
		SetTimeout(10 * time.Second).
		SetAuthToken(authToken).
		SetAuthScheme("Bot")
	m.client.SetBaseURL(domain)
}

func (m *MyHttpClient) GetMethod(method string) string {
	// 发送请求
	resp, err := m.client.R().
		Get(method)

	// 检查错误
	if err != nil {
		log.Fatalf("请求错误: %v", err)
	}
	if resp.StatusCode() == 200 {
		var response SocketIpStruct
		err := json.Unmarshal(resp.Body(), &response)
		if err != nil {
			log.Fatalln(err)
		}

		return response.URL
	}
	return ""
}

func (m *MyHttpClient) PostMethod(method string, body interface{}) string {
	// 发送请求
	_, err := m.client.R().
		SetBody(body).
		Post(method)
	// 检查错误
	if err != nil {
		log.Fatalf("请求错误: %v", err)
	}
	return ""
}

func (m *MyHttpClient) PostMethodParam(method string, paramKey string, paramValue string, body interface{}) string {
	// 发送请求
	_, err := m.client.R().
		SetPathParam(paramKey, paramValue).
		SetBody(body).
		Post(method)
	// 检查错误
	if err != nil {
		log.Fatalf("请求错误: %v", err)
	}
	return ""
}
