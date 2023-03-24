package service

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"m800/internal/pkg/model"
	"m800/internal/pkg/model/dto"
)

type LineMessageService struct {
	repo *dto.LineMessageRepository
}

func NewLineMessageService() *LineMessageService {

	return &LineMessageService{
		repo: dto.NewLineMessageRepository(),
	}
}

func (l *LineMessageService) CreateMessage(messageId string, messageType model.MessageType, text string, source *linebot.EventSource) (*model.LineMessage, error) {

	res, err := l.repo.Create(messageId, string(messageType), text, source)
	if err != nil {
		return nil, err
	}

	message := &model.LineMessage{
		Id:          res.Id,
		MessageType: model.MessageType(res.MessageType),
		Text:        res.Text,
		Source:      res.Source,
	}

	return message, nil
}
func (l *LineMessageService) GetAllMessagesByUser(id string) ([]*model.LineMessage, error) {
	res, err := l.repo.GetAllMessagesByUser(id)
	if err != nil {
		return nil, err
	}

	messages := []*model.LineMessage{}
	for _, message := range res {
		messages = append(messages, &model.LineMessage{
			Id:          message.Id,
			MessageType: model.MessageType(message.MessageType),
			Text:        message.Text,
			Source:      message.Source,
		})
	}

	return messages, nil
}

func (l *LineMessageService) GetMessageByID(id string) (*model.LineMessage, error) {

	res, err := l.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	message := &model.LineMessage{
		Id:          res.Id,
		MessageType: model.MessageType(res.MessageType),
		Text:        res.Text,
	}

	return message, nil
}
