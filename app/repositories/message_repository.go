package repositories

import (
	"bitbucket.org/frchandra/giscust/app/models"
	"bitbucket.org/frchandra/giscust/app/validations"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (this *MessageRepository) SearchMessagesByRoomId(roomId string) (models.Message, error) {
	var messageFound models.Message
	err := this.db.Where("room_id = ?", roomId).First(&messageFound).Error
	return messageFound, err
}

func (this *MessageRepository) UpsertMessages(message *validations.Message) (models.Message, error) {
	messageFound, err := this.SearchMessagesByRoomId(message.RoomId)

	messageFound.AppId = message.AppId
	messageFound.AvatarUrl = message.AvatarUrl
	messageFound.Email = message.Email
	messageFound.IsNewSession = message.IsNewSession
	messageFound.IsResolved = message.IsResolved
	messageFound.Name = message.Name
	messageFound.RoomId = message.RoomId
	messageFound.Source = message.Source

	this.db.Save(&messageFound)
	return messageFound, err
}

func (this *MessageRepository) UpdateMessageHandled(roomId string) error {
	err := this.db.Model(&models.Message{}).Where("room_id = ?", roomId).Update("is_handled", true).Error
	return err
}
