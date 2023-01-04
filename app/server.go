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
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) Initialize(appConfig *config.AppConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	//Try to connect with database
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", appConfig.DBHost, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName, appConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed on connecting to the database server")
	} else {
		fmt.Println("Database connection established")
		fmt.Println("Using database " + server.DB.Migrator().CurrentDatabase())
	}

	if appConfig.IsProduction == "false" {
		gin.SetMode(gin.DebugMode)
	}

	server.Router = gin.Default()

	server.initializeRoutes(server.Router)

	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database migrated successfully")

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
