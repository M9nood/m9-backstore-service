package message

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineMessagePkg struct {
	Data interface{}
}

type LineMessageInterface interface {
	DetailFlexMessage() *linebot.FlexMessage
}

func InitLineMessagePkg(data interface{}) LineMessageInterface {
	return &LineMessagePkg{Data: data}
}
