package http

import (
	controller "m9-backstore-service/controllers"

	"github.com/labstack/echo/v4"
)

type lineBotHttpRoute struct {
	handler controller.BotHandler
}

func NewLineBotHttpRoute(botHandler controller.BotHandler) lineBotHttpRoute {
	return lineBotHttpRoute{
		handler: botHandler,
	}
}

func (h lineBotHttpRoute) Route(e *echo.Echo) {
	botRoute := e.Group("/line/bot")
	botRoute.POST("/webhook", h.handler.BotReplyHandler)
}
