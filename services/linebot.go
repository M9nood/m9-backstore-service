package service

import (
	"fmt"
	"log"
	"m9-backstore-service/constant"
	"m9-backstore-service/models/line"
	"m9-backstore-service/models/product"
	linepkg "m9-backstore-service/pkg/line"
	"os"

	repository "m9-backstore-service/repositories"

	"github.com/jinzhu/gorm"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineBotService struct {
	ChannelSecret       string
	ChannelAccesssToken string
	Db                  *gorm.DB
	UserRepo            repository.UserReposityInterface
}

type LineBotServiceInterface interface {
	InitBot() (bot *linebot.Client)
	WatchAndReplyMessage(lineMsg *line.LineMessage) error
	CreateMessageByTriggerMessage(botBrain line.BotBrain) ([]linebot.SendingMessage, error)
}

func NewLineBotService(db *gorm.DB) LineBotServiceInterface {
	userRepo := repository.NewUserReposity(db)
	return &LineBotService{
		ChannelSecret:       os.Getenv("LINE_CHANNEL_SECRET"),
		ChannelAccesssToken: os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
		Db:                  db,
		UserRepo:            userRepo,
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
		botBrain := linepkg.CreateBotBrain(lineMsg.Events[0])
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
	switch triggerMsg := botBrain.Action; triggerMsg {
	case constant.Greeting:
		newMessages, _ = s.MessageCallingBot(botBrain)
	case constant.GetProduct:
		if botBrain.Code != "" {
			newMessages, _ = s.MessageGetProduct(botBrain)
		} else {
			newMessages, _ = s.MessageGetProducts(botBrain)
		}
	default:
		newMessages, _ = s.MessageCallingBot(botBrain)
	}
	return newMessages, nil
}

func (s LineBotService) MessageGetProducts(botBrain line.BotBrain) ([]linebot.SendingMessage, error) {
	var newMessages []linebot.SendingMessage
	var replyMsg string
	productRepo := repository.NewProductReposity(s.Db)
	key := 1
	result, err := productRepo.GetProducts(&key, product.ProductQueryParams{})
	if err != nil {
		return newMessages, err
	}
	replyMsg += fmt.Sprintf("%s :\n", botBrain.Title)
	for _, item := range result.Products {
		replyMsg += product.ToLineMessage(item)
	}
	newMessages = append(newMessages, linebot.NewTextMessage(replyMsg))
	return newMessages, nil
}

func (s LineBotService) MessageGetProduct(botBrain line.BotBrain) ([]linebot.SendingMessage, error) {
	var newMessages []linebot.SendingMessage
	var replyMsg string
	productRepo := repository.NewProductReposity(s.Db)
	result, err := productRepo.GetProductByCode(botBrain.Code)
	if err != nil {
		if err.GetCode() == "404" {
			newMessages = append(newMessages, linebot.NewTextMessage(err.Error()))
			return newMessages, nil
		}
		return newMessages, err
	}
	replyMsg += fmt.Sprintf("%s %s", result.DispCode, result.ProductName)

	newMessages = append(newMessages, linebot.NewTextMessage(replyMsg))
	return newMessages, nil
}

func (s LineBotService) MessageCallingBot(botBrain line.BotBrain) ([]linebot.SendingMessage, error) {
	var newMessages []linebot.SendingMessage
	message := linebot.NewTextMessage("Hello! Can I help you").
		WithQuickReplies(
			linebot.NewQuickReplyItems(
				linebot.NewQuickReplyButton("", linebot.NewMessageAction("#view", "#view")),
			),
		)
	newMessages = append(newMessages, message)
	return newMessages, nil
}
