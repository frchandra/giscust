package services

import (
	"bitbucket.org/frchandra/giscust/app/models"
	"bitbucket.org/frchandra/giscust/app/repositories"
	"bitbucket.org/frchandra/giscust/app/validations"
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
		if strings.Contains(agent.Name, "chandra") {
			myAgents = append(myAgents, agent)
		}
	}
	return myAgents
}

func (this *MessagesService) GetOnlyAvailableAgent(agents []validations.Agent) validations.Agent {
	rand.Shuffle(len(agents), func(i, j int) {
		agents[i], agents[j] = agents[j], agents[i]
	})

	sort.Slice(agents, func(i, j int) bool {
		return agents[i].Id < agents[j].Id
	})

	var availableAgents []validations.Agent
	for _, agent := range agents {
		if agent.CurrentCustomerCount < 2 {
			availableAgents = append(availableAgents, agent)
		}
	}

	return agents[0]
}
