package service

import (
	"fmt"
	"log"
	"m9-backstore-service/constant"
	"m9-backstore-service/models/line"
	"m9-backstore-service/models/product"
	util "m9-backstore-service/utils"
	"os"

	repository "m9-backstore-service/repositories"

	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineBotService struct {
	ChannelSecret       string
	ChannelAccesssToken string
	Db                  *gorm.DB
}

type LineBotServiceInterface interface {
	InitBot() (bot *linebot.Client)
	WatchAndReplyMessage(lineMsg *line.LineMessage) error
	CreateMessageByTriggerMessage(botBrain line.BotBrain) ([]linebot.SendingMessage, error)
}

func NewLineBotService(db *gorm.DB) LineBotServiceInterface {
	return &LineBotService{
		ChannelSecret:       os.Getenv("LINE_CHANNEL_SECRET"),
		ChannelAccesssToken: os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
		Db:                  db,
	}
}

func (s LineBotService) InitBot() (bot *linebot.Client) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Bot not working", r)
			bot = nil
		}
	}()
	bot, err := linebot.New(
		s.ChannelSecret,
		s.ChannelAccesssToken,
	)
	if err != nil {
		panic("error init bot")
	}
	return bot
}

func (s LineBotService) WatchAndReplyMessage(lineMsg *line.LineMessage) error {
	bot := s.InitBot()
	if bot == nil {
		return nil
	}
	if len(lineMsg.Events) > 0 {
		botBrain := util.CreateBotBrain(lineMsg.Events[0])
		if botBrain.Action != "" {
			msg, err := s.CreateMessageByTriggerMessage(botBrain)
			if err != nil {
				return nil
			}
			if _, err = bot.ReplyMessage(lineMsg.Events[0].ReplyToken, msg...).Do(); err != nil {
				log.Print(err)
			}
		}
	}
	return nil
}

func (s LineBotService) CreateMessageByTriggerMessage(botBrain line.BotBrain) ([]linebot.SendingMessage, error) {
	var newMessages []linebot.SendingMessage
	var replyMsg string
	switch triggerMsg := botBrain.Action; triggerMsg {
	case constant.GetProduct:
		productRepo := repository.NewProductReposity(s.Db)
		result, err := productRepo.GetProducts()
		if err != nil {
			fmt.Println("err", err)
		}
		replyMsg += fmt.Sprintf("%s :\n", botBrain.Title)
		for _, item := range result {
			replyMsg += product.ToLineMessage(item)
		}
		newMessages = append(newMessages, linebot.NewTextMessage(replyMsg))
		return newMessages, nil
	default:
		break
	}
	return newMessages, nil
}
