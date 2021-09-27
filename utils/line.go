package util

import (
	"m9-backstore-service/constant"
	"m9-backstore-service/models/line"
	"strings"
)

func CreateBotBrain(event line.LineMessageEvent) line.BotBrain {
	bb := line.BotBrain{}
	switch msgType := event.Message.Type; msgType {
	case "text":
		msgArg := strings.Split(strings.TrimSpace(strings.ToLower(event.Message.Text)), " ")
		if len(msgArg) == 1 {
			if _, ok := constant.BotDictionaries[msgArg[0]]; ok {
				bb = constant.BotDictionaries[msgArg[0]]
				bb.InputMessage = event.Message
			}
		} else if len(msgArg) == 2 {
			if _, ok := constant.BotDictionaries[msgArg[0]]; ok {
				bb = constant.BotDictionaries[msgArg[0]]
				bb.InputMessage = event.Message
				bb.Code = msgArg[1]
				bb.ReplyMessage.ReplyToken = event.ReplyToken
			}
		}
	}
	return bb
}
