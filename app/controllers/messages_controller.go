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
	if err = c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":         "fail",
			"error_messages": err.Error(),
		})
		return
	}

	/*	var messageFound models.Message
		if messageFound, err = this.messagesService.UpsertMessages(&message); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "creating new messages",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "updating previous message",
			"data":    messageFound,
		})
		return*/

	agentsListResponse, err := utils.GetAllAgentsByDivision()
	agents := agentsListResponse.Data
	agents = this.messagesService.GetOnlyMyAgents(agents)
	if len(agents) < 1 {
		c.JSON(http.StatusAccepted, gin.H{
			"status":  "success, queued",
			"message": "this message is temporarily on queue",
		})
	}
	agents = this.messagesService.GetOnlyAvailableAgents(agents)

	//assign agent to a room

	//update messages status

	//return response

}
