package service

import (
	"fitus-chat-service/internal/model"
	"fitus-chat-service/internal/repository"
)

type MessageService struct {
	repo *repository.MessageRepository
}

func NewMessageService(repo *repository.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(user uint, content string) error {
	message := &model.Message{
		UserId:    user,
		Content: content,
	}
	return s.repo.Create(message)
}

func (s *MessageService) GetMessageByID(id uint) (*model.Message, error) {
	return s.repo.FindByID(id)
}

func (s *MessageService) GetAllMessages() ([]model.Message, error) {
	return s.repo.FindAll()
}

func (s *MessageService) UpdateMessage(message *model.Message) error {
	return s.repo.Update(message)
}

func (s *MessageService) DeleteMessage(id uint) error {
	return s.repo.Delete(id)
}
