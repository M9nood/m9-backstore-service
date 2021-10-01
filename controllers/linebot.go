package controller

import (
	"log"
	"net/http"

	"m9-backstore-service/models/line"
	service "m9-backstore-service/services"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type BotHandler struct {
	lineBotService service.LineBotServiceInterface
}

func NewLineBotController(db *gorm.DB) BotHandler {
	svc := service.NewLineBotService(db)
	return BotHandler{
		lineBotService: svc,
	}
}

func (h *BotHandler) BotReplyHandler(c echo.Context) error {
	Line := new(line.LineMessage)
	if err := c.Bind(Line); err != nil {
		log.Println("err")
		return c.String(http.StatusInternalServerError, "error")
	}
	if len(Line.Events) > 0 {
		h.lineBotService.WatchAndReplyMessage(Line)
	}
	return c.String(http.StatusOK, "ok")
}
