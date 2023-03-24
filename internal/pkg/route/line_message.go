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

type LineMessageSendingRequest struct {
	Message  string `json:"message"`
	TargetId string `json:"target_id"`
}

func SetUpLineMessage(r *gin.Engine) {
	r.POST("/line_message/callback", func(c *gin.Context) {
		lineMessageCallback(c)
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.POST("/line_message/send", func(c *gin.Context) {
		req := &LineMessageSendingRequest{}
		err := c.Bind(req)
		if err != nil {
			log.Default().Println(err)
			c.String(http.StatusBadRequest, "Invalid request body")
			return
		}

		message := linebot.NewTextMessage(req.Message)

		bot := line.NewBot()
		res, err := bot.PushMessage(req.TargetId, message)
		if err != nil {
			c.String(http.StatusBadRequest, "Send message error: %v", err)
		} else {
			c.JSON(http.StatusOK, res)
		}
	})

	// api to get all messages by user id
	r.GET("/line_message/:user_id", func(c *gin.Context) {
		user_id := c.Param("user_id")
		s := service.NewLineMessageService()
		messages, err := s.GetAllMessagesByUser(user_id)
		if err != nil {
			c.String(http.StatusBadRequest, "Get messages error: %v", err)
		} else {
			c.JSON(http.StatusOK, messages)
		}
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
