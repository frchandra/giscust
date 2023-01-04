package app

import (
	"bitbucket.org/frchandra/giscust/app/models"
	"bitbucket.org/frchandra/giscust/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gormigrate/gormigrate/v2"
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

	// Running migration
	/*	for _, model := range RegisterModels() {
			err = server.DB.Debug().AutoMigrate(model.Model)
		}
		if err != nil {
			log.Fatal(err)
		}*/

	m := gormigrate.New(server.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "201608301400",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(models.Message{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("messages")
			},
		},
	})

	err = m.Migrate()
	if err == nil {
		fmt.Println("Migration did run successfully")
	} else {
		fmt.Println("Could not migrate: %v", err)
	}

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
