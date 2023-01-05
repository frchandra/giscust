package controllers

import (
	"bitbucket.org/frchandra/giscust/app/models"
	"bitbucket.org/frchandra/giscust/app/services"
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
	err = c.ShouldBindJSON(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":         "fail",
			"error_messages": err.Error(),
		})
		return
	}

	var messageFound models.Message
	messageFound, err = this.messagesService.UpsertMessages(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":         "fail",
			"error_messages": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   messageFound,
	})

}
