package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"m800/internal/pkg/line"
	"m800/internal/pkg/model"
	"m800/internal/pkg/service"
	"net/http"
)

func SetUpLineMessage(r *gin.Engine) {
	r.POST("/line_message/callback", func(c *gin.Context) {
		lineMessageCallback(c)
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

}

func lineMessageCallback(c *gin.Context) {
	bot := line.NewBot()
	events, err := bot.ParseRequest(c.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.String(http.StatusBadRequest, err.Error())
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {

			case *linebot.TextMessage:
				s := service.NewLineMessageService()
				fmt.Println(event.Source.UserID)
				_, err := s.CreateMessage(message.ID, model.TextMessageType, message.Text, event.Source)
				if err != nil {
					log.Default().Println(err)
				}
			}
		}
	}
}
