package app

import (
	"bitbucket.org/frchandra/giscust/app/controllers"
	"bitbucket.org/frchandra/giscust/app/repositories"
	"bitbucket.org/frchandra/giscust/app/services"
)

func (server *Server) initializeRoutes() {

	//Manual dependency injection. This is not good, still need improved
	messageRepo := repositories.NewRepository(server.Database)
	messageService := services.NewMessageService(messageRepo)
	messageController := controllers.NewMessageController(messageService)

	v1 := server.Router.Group("/api/v1")
	v1.GET("/", controllers.Index)
	v1.POST("/new_request", messageController.HandleMessages)
	v1.POST("/close_request", messageController.HandleSettlement)

}
