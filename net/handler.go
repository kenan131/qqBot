package net

import (
	"helloGo/dto"
	"helloGo/log"
)

var eventHandlerMap = map[dto.OPCode]map[dto.EventType]EventHandler{}

// EventHandler 抽象方法
type EventHandler func(event *dto.WSPayload, message []byte) error

func RegisterHandler(code dto.OPCode, eventT dto.EventType, handler EventHandler) {
	// 添加到map中
	if _, exists := eventHandlerMap[code]; !exists {
		eventHandlerMap[code] = make(map[dto.EventType]EventHandler)
	}
	eventHandlerMap[code][eventT] = handler
}

func HandlerProcess(code dto.OPCode, eventT dto.EventType, payload *dto.WSPayload) error {
	if tempMap, ok := eventHandlerMap[code]; ok {
		if handler, ok1 := tempMap[eventT]; ok1 {
			// 调用 eventHandler
			handler(payload, payload.RawMessage)
		} else {
			log.Infof("没有添加该事件的处理器,opCode:%d eventType:%s", code, eventT)
		}
	}
	return nil
}
