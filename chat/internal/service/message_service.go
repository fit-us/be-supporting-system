package service

import (
	"fitus-chat-service/internal/model"
	"fitus-chat-service/internal/repository"
	"fmt"

	"github.com/gorilla/websocket"
)

type MessageService struct {
	repo *repository.MessageRepository
}

var (
	wsClients = make(map[*websocket.Conn]bool)
	broadcast = make(chan model.Message)
)

func init() {
	go handleMessages()
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range wsClients {
			if err := client.WriteJSON(msg); err != nil {
				fmt.Println("Broadcast error:", err)
				client.Close()
				delete(wsClients, client)
			}
		}
	}
}

func RegisterClient(client *websocket.Conn) {
	wsClients[client] = true
}

func UnregisterClient(client *websocket.Conn) {
	delete(wsClients, client)
}

func NewMessageService(repo *repository.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(user uint, content string) error {
	message := &model.Message{
		UserId:  user,
		Content: content,
	}
	if err := s.repo.Create(message); err != nil {
		return err
	}
	broadcast <- *message
	return nil
}

func (s *MessageService) GetMessageByID(id uint) (*model.Message, error) {
	return s.repo.FindByID(id)
}

func (s *MessageService) UpdateMessage(message *model.Message) error {
	return s.repo.Update(message)
}

func (s *MessageService) DeleteMessage(id uint) error {
	return s.repo.Delete(id)
}
