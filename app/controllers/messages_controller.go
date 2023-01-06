package controllers

import (
	"bitbucket.org/frchandra/giscust/app/services"
	"bitbucket.org/frchandra/giscust/app/utils"
	"bitbucket.org/frchandra/giscust/app/validations"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessagesController struct {
	messagesService *services.MessagesService
}

func NewMessageController(ms *services.MessagesService) *MessagesController {
	return &MessagesController{ms}
}

func (this *MessagesController) HandleMessages(c *gin.Context) {
	var message validations.Message
	var err error

	//Decode message body
	if err = c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":         "fail",
			"error_messages": err.Error(),
		})
		return
	}

	//Upsert the new coming message to db
	_, err = this.messagesService.UpsertMessages(&message)

	//Get all agent
	agentsListResponse, err := utils.GetAllAgentsByDivision()
	agents := agentsListResponse.Data
	agents = this.messagesService.GetOnlyMyAgents(agents)
	agent, err := this.messagesService.GetOnlyAvailableAgent(agents)

	//If there is no appropriate agent, send a clarifying response. Message already queued on db
	if agent == nil || err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"status":  "success, queued",
			"message": "this message is temporarily on queue",
			"details": err,
		})
		return
	}

	//assign agent to a room
	//response := utils.AssignAgentToRoom(agent[0].Id, message.RoomId)

	//update messages status
	this.messagesService.UpdateMessageHandled(message.RoomId)

	//return response
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		//"qiscus_agent_allocator_response": string(response),
		"roomId": message.RoomId,
		"agenId": agents[0].Id,
	})
}
