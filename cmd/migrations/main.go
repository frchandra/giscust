package main

import (
	"bitbucket.org/frchandra/giscust/config"
	"bitbucket.org/frchandra/giscust/database/migrations"
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	fmt.Println("hello world")
	RunMigration(os.Args[1], os.Args[2])
}

func RunMigration(option string, id string) error {
	appConfig := config.GetAppConfig()

	var err error
	var database *gorm.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", appConfig.DBHost, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName, appConfig.DBPort)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed on connecting to the database server")
	} else {
		fmt.Println("Database connection established")
		fmt.Println("Using database " + database.Migrator().CurrentDatabase())
	}

	var migrationsData []*gormigrate.Migration = migrations.GetAllMigrations()
	m := gormigrate.New(database, gormigrate.DefaultOptions, migrationsData)

	fmt.Println("option " + option + " is chosen")
	fmt.Println("id " + id + " is chosen")

	switch option {
	case "m":
		err = m.Migrate()
	case "f":
		err = m.RollbackTo("init")
	case "mf":
		err = m.RollbackTo("init")
		if err == nil {
			err = m.Migrate()
		}
	case "mfs":
		err = m.RollbackTo("init")
		if err != nil {
			err = m.Migrate()
		}
	case "rollbackto":
		err = m.RollbackTo(id)
	case "migrateto":
		err = m.MigrateTo(id)
	default:
		panic("option " + option + " unknown")
	}

	//err = m.Migrate()

	if err == nil {
		fmt.Println("Migration did run successfully")
	} else {
		fmt.Println("Could not migrate: %v", err)
	}

	return err
}
