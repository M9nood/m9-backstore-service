package constant

import (
	"m9-backstore-service/models/line"
)

var (
	GetProduct = "GetProduct"
)

var BotDictionaries = map[string]line.BotBrain{
	"view": {
		InputMessage: line.Message{},
		ReplyMessage: line.ReplyMessage{},
		Title:        "Product List",
		Action:       GetProduct,
		Code:         "",
	},
}
