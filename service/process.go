package service

import (
	"database/sql"
	"helloGo/dto"
	"helloGo/net"
)

type Processor struct {
	Api net.MyHttpClient
	Db  *sql.DB
}

func (p *Processor) ProcessAtMessage(data *dto.Message) error {
	match := dto.CommandRegular.FindStringSubmatch(data.Content)
	var content = ""
	if len(match) > 2 {
		switch match[1] {
		case "001":
			content = "001指令"
		case "002":
			content = "002指令"
		case "003":

		}
	} else {
		content = "固定回复！"
	}
	replyMessage := &dto.MessageToCreate{
		Content: content,
		MessageReference: &dto.MessageReference{
			// 引用这条消息
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	}
	p.Api.PostMethodParam(dto.MessagesURI, "channel_id", data.ChannelID, replyMessage)
	return nil
}

func (p *Processor) ProcessMessage(data *dto.Message) error {
	if replyContent, exists := dto.DefaultMessage[data.Content]; exists {
		replyMessage := &dto.MessageToCreate{
			Content: replyContent,
			MessageReference: &dto.MessageReference{
				// 引用这条消息
				MessageID:             data.ID,
				IgnoreGetMessageError: true,
			},
		}
		p.Api.PostMethodParam(dto.MessagesURI, "channel_id", data.ChannelID, replyMessage)
	}

	return nil
}
