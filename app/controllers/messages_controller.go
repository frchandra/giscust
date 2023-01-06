package controllers

import (
	"bitbucket.org/frchandra/giscust/app/models"
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
	response := utils.AssignAgentToRoom(agent[0].Id, message.RoomId)

	//update messages status
	this.messagesService.UpdateMessageHandled(message.RoomId)

	//return response
	c.JSON(http.StatusOK, gin.H{
		"status":                          "success",
		"qiscus_agent_allocator_response": string(response),
		"roomId":                          message.RoomId,
		"agenId":                          agents[0].Id,
	})
}

func (this *MessagesController) HandleSettlement(c *gin.Context) {
	var settlement validations.Settlement
	var err error

	//Decode message body
	if err = c.ShouldBindJSON(&settlement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":         "fail",
			"error_messages": err.Error(),
		})
		return
	}

	//Set the is_resolved field on the settled message to true
	err = this.messagesService.UpdateMessageSettled(settlement.Service.RoomId)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"status":  "success, with exception",
			"message": "did not found the issued message",
		})
	}

	//Get the oldest unresolved message
	var message models.Message
	message, err = this.messagesService.GetOldestUnresolved()

	//If there is no message left, send success response. The application standby for a new message
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"status":  "success, with exception",
			"message": "did not found the issued message",
		})
		return
	}

	//If there is a message, check if there is available agent
	//Get all agent
	agentsListResponse, err := utils.GetAllAgentsByDivision()
	agents := agentsListResponse.Data
	agents = this.messagesService.GetOnlyMyAgents(agents)
	agent, err := this.messagesService.GetOnlyAvailableAgent(agents)

	//If there is no appropriate agent, send a clarifying response. Message already queued on db
	if agent == nil || err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"status":  "success, queued",
			"message": "the next message is temporarily on queue",
			"details": err,
		})
		return
	}

	//If there is an available agent, allocate the corresponding agent to the next message
	response := utils.AssignAgentToRoom(agent[0].Id, message.RoomId)

	//update messages status
	this.messagesService.UpdateMessageHandled(message.RoomId)

	//send response
	c.JSON(http.StatusOK, gin.H{
		"status":                          "success",
		"qiscus_agent_allocator_response": string(response),
		"roomId":                          message.RoomId,
		"agenId":                          agents[0].Id,
	})
}
