package migrations

import (
	"bitbucket.org/frchandra/giscust/app/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func GetAllMigrations() []*gormigrate.Migration {
	var migrations = []*gormigrate.Migration{
		{
			ID: "init",
			Migrate: func(tx *gorm.DB) error {
				return tx.Debug().AutoMigrate()
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Debug().Migrator().DropTable()
			},
		},
		{
			ID: "create_person_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.Debug().AutoMigrate(models.Message{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Debug().Migrator().DropTable(models.Message{})
			},
		},
	}
	return migrations
}
