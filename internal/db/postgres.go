package db

import (
	"api/internal/db/domain"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	log.Info().Msg("connecting to database...")

	dsn := viper.GetString("database.dsn")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: createGormLogger(),
	})
	if err != nil {
		log.Error().Err(err).Msg("could not connect to postgres")
		return nil, err
	}

	err = db.AutoMigrate(&domain.ShortLink{})

	if err != nil {
		log.Error().Err(err).Msg("could not migrate database")
		return nil, err
	}

	return db, nil
}
