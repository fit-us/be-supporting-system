package repository

import (
	"fitus-chat-service/internal/model"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(message *model.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepository) FindByID(id uint) (*model.Message, error) {
	var message model.Message
	err := r.db.First(&message, id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *MessageRepository) FindAll() ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *MessageRepository) Update(message *model.Message) error {
	return r.db.Save(message).Error
}

func (r *MessageRepository) Delete(id uint) error {
	return r.db.Delete(&model.Message{}, id).Error
}