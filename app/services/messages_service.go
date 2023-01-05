package services

import (
	"bitbucket.org/frchandra/giscust/app/models"
	"bitbucket.org/frchandra/giscust/app/repositories"
	"bitbucket.org/frchandra/giscust/app/validations"
)

type MessagesService struct {
	messageRepository *repositories.MessageRepository
}

func NewMessageService(mr *repositories.MessageRepository) *MessagesService {
	return &MessagesService{mr}
}

func (this *MessagesService) UpsertMessages(message *validations.Message) (models.Message, error) {
	return this.messageRepository.UpsertMessages(message)
}
