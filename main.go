package main

import (
	"api/cmd"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
	}

	// Config options and paths
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/opt/short-me")

	// Set defaults and bindings
	viper.SetDefault("debug", false)
	_ = viper.BindEnv("database.dsn", "DATABASE_DSN")

	// Load config values from environment when provided
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("no config file found")
			os.Exit(1)
		} else {
			fmt.Println("could not load config")
			os.Exit(1)
		}
	}

	// Initialize logger
	debugMode := viper.GetBool("debug")
	if debugMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)

	} else {
		zerolog.TimestampFieldName = "ts"
		zerolog.LevelFieldName = "level"
		zerolog.MessageFieldName = "message"
		log.Logger = log.Output(os.Stderr).Level(zerolog.InfoLevel)
	}

	// Watch for config changes / auto reloading
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info().Str("config", e.Name).Msg("reloading changed config file")
	})
	viper.WatchConfig()

	cmd.Execute()
}
