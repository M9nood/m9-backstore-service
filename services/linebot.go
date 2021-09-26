package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"m9-backstore-service/constant"
	"m9-backstore-service/models/line"
	"m9-backstore-service/models/product"
	"net/http"
	"os"

	repository "m9-backstore-service/repositories"

	"github.com/jinzhu/gorm"
)

type LineBotService struct {
	ChannelSecret       string
	ChannelAccesssToken string
	Db                  *gorm.DB
}

var lineBotServiceInstance *LineBotService

func NewLineBotService(db *gorm.DB) *LineBotService {
	if lineBotServiceInstance == nil {
		lineBotServiceInstance = &LineBotService{
			ChannelSecret:       os.Getenv("LINE_CHANNEL_SECRET"),
			ChannelAccesssToken: os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
			Db:                  db,
		}
	}
	return lineBotServiceInstance
}

func (s LineBotService) WatchAndReplyMessage(lineMsg *line.LineMessage) error {
	if len(lineMsg.Events) > 0 {
		message := lineMsg.Events[0].Message.Text
		if _, ok := constant.BotWords[message]; ok {
			msg, err := s.createMessageByTriggerMessage(message)
			if err != nil {
				return nil
			}
			message := line.ReplyMessage{
				ReplyToken: lineMsg.Events[0].ReplyToken,
				Messages:   msg,
			}
			go s.replyMessageLine(message)
		}
		return nil
	}
	return nil
}

func (s LineBotService) replyMessageLine(Message line.ReplyMessage) error {
	value, _ := json.Marshal(Message)

	url := "https://api.line.me/v2/bot/message/reply"

	var jsonStr = []byte(value)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.ChannelAccesssToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))

	return err
}

func (s LineBotService) createMessageByTriggerMessage(message string) ([]line.Text, error) {
	text := []line.Text{}
	var replyMsg string
	switch triggerMsg := constant.BotWords[message]; triggerMsg {
	case constant.BotWords["view"]:
		productRepo := repository.NewProductReposity(s.Db)
		result, err := productRepo.GetProducts()
		if err != nil {
			fmt.Println("err", err)
		}
		for _, item := range result {
			replyMsg += product.ToLineMessage(item)
		}

		text = append(text, line.Text{Type: "text", Text: replyMsg})
		return text, nil
	default:
		break
	}
	return text, nil
}
