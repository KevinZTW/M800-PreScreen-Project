package dto

import (
	"context"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LineMessage struct {
	Id          string
	MessageType string `json:"message_type" bson:"message_type"`
	Text        string
	Source      *linebot.EventSource
}

type LineMessageRepository struct {
	collection *mongo.Collection
}

const LineMessageCollection = "LineMessageCollection"

func NewLineMessageRepository() *LineMessageRepository {
	db := MongoDB()
	return &LineMessageRepository{
		collection: db.Collection(LineMessageCollection),
	}
}

func (l *LineMessageRepository) Create(id string, messageType string, text string, source *linebot.EventSource) (*LineMessage, error) {
	message := &LineMessage{
		Id:          id,
		MessageType: messageType,
		Text:        text,
		Source:      source,
	}
	_, err := l.collection.InsertOne(context.TODO(), message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (l *LineMessageRepository) GetByID(id string) (*LineMessage, error) {
	message := &LineMessage{}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = l.collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}
