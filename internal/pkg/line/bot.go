package line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"m800/internal/pkg/config"
	"net/http"
	"sync"
)

type Bot struct {
	Client *linebot.Client
}

var (
	once sync.Once
	bot  *Bot
)

func NewBot() *Bot {
	once.Do(func() {
		con := config.New()
		client, err := linebot.New(con.LineMessageAPI.ChannelSecrete, con.LineMessageAPI.ChannelAccessToken)
		if err != nil {
			log.Fatal(err)
		}
		res, err := client.GetBotInfo().Do()
		if err != nil || res == nil {
			log.Fatal(err)
		}
		log.Printf("Create new bot client: %s", res.DisplayName)
		bot = &Bot{Client: client}
	})

	return bot
}

func (b *Bot) ParseRequest(r *http.Request) ([]*linebot.Event, error) {
	return b.Client.ParseRequest(r)
}

func (b *Bot) PushMessage(to string, messages ...linebot.SendingMessage) (*linebot.BasicResponse, error) {
	return b.Client.PushMessage(to, messages...).Do()
}
