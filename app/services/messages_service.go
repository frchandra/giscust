package services

import (
	"bitbucket.org/frchandra/giscust/app/models"
	"bitbucket.org/frchandra/giscust/app/repositories"
	"bitbucket.org/frchandra/giscust/app/validations"
	"errors"
	"math/rand"
	"sort"
	"strings"
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

func (this *MessagesService) GetOnlyMyAgents(agents []validations.Agent) []validations.Agent {
	var myAgents []validations.Agent
	for _, agent := range agents {
		if strings.Contains(agent.Name, "chandra1") {
			myAgents = append(myAgents, agent)
		}
	}
	return myAgents
}

func (this *MessagesService) GetOnlyAvailableAgent(agents []validations.Agent) ([]validations.Agent, error) {
	//Shuffle the available agent list
	rand.Shuffle(len(agents), func(i, j int) {
		agents[i], agents[j] = agents[j], agents[i]
	})

	//Sort the agent. The agent with the less current customer count gets on top
	sort.Slice(agents, func(i, j int) bool {
		return agents[i].CurrentCustomerCount < agents[j].CurrentCustomerCount
	})

	//Apply custom requirement: agent should only handle maximum 2 customer
	var availableAgents []validations.Agent
	for _, agent := range agents {
		if agent.CurrentCustomerCount < 2 {
			availableAgents = append(availableAgents, agent)
		}
	}

	var err error
	if len(availableAgents) < 1 {
		err = errors.New("not found available agents")
	}
	return availableAgents, err

}

func (this *MessagesService) UpdateMessageHandled(roomId string) error {
	return this.messageRepository.UpdateMessageHandled(roomId)
}

func (this *MessagesService) UpdateMessageSettled(roomId string) error {
	return this.messageRepository.UpdateMessageSettled(roomId)
}

func (this *MessagesService) GetOldestUnresolved() (models.Message, error) {
	return this.messageRepository.GetOldestUnresolved()
}
