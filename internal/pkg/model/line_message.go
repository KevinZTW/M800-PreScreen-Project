package model

import "github.com/line/line-bot-sdk-go/v7/linebot"

const (
	TextMessageType MessageType = "TextMessageType"
)

type MessageType string

type LineMessage struct {
	Id          string               `json:"id" bson:"id"`
	MessageType MessageType          `json:"message_type" bson:"message_type"`
	Text        string               `json:"text" bson:"text"`
	Source      *linebot.EventSource `json:"source" bson:"source"`
}
