package app

import (
	"bitbucket.org/frchandra/giscust/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	Database *gorm.DB
	Router   *gin.Engine
}

func (server *Server) initializeDb(appConfig *config.AppConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", appConfig.DBHost, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName, appConfig.DBPort)
	server.Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed on connecting to the database server")
	} else {
		fmt.Println("Database connection established")
		fmt.Println("Using database " + server.Database.Migrator().CurrentDatabase())
	}
}

func (server *Server) Initialize(appConfig *config.AppConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.initializeDb(appConfig)

	if appConfig.IsProduction == "false" {
		gin.SetMode(gin.DebugMode)
	}

	server.Router = gin.Default()

	server.initializeRoutes()

}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	err := server.Router.Run(addr)
	if err != nil {
		log.Fatal("Server unable to start")
	}
}

func Run() {
	appConfig := config.GetAppConfig()
	var server = Server{}
	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}
