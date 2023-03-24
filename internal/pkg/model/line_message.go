package model

import "github.com/line/line-bot-sdk-go/v7/linebot"

const (
	TextMessageType MessageType = "TextMessageType"
)

type MessageType string

type LineMessage struct {
	Id          string
	MessageType MessageType
	Text        string
	Source      *linebot.EventSource
}
