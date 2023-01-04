package app

import (
	"bitbucket.org/frchandra/giscust/app/controllers"
	"github.com/gin-gonic/gin"
)

func (server *Server) initializeRoutes(router *gin.Engine) {
	//server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	v1 := router.Group("/v1")
	v1.GET("/", controllers.Index)

}
