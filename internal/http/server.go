package http

import (
	"api/internal/config"
	"api/internal/db"
	"api/internal/http/handlers"
	"api/internal/http/middleware"
	"api/internal/http/services"
	"api/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// Serve godoc
// Serves the API server on a predefined port
func Serve() {
	log.Info().Msg("starting server...")

	// Database connection
	postgres, err := db.InitPostgres()

	if err != nil {
		os.Exit(1)
	}

	// Create fiber app
	app := fiber.New(config.GetFiberConfig())

	// Global middleware
	app.Use(requestid.New(requestid.Config{
		Generator: func() string {
			reqid := uuid.New().String()
			log.Logger = log.Logger.With().Interface("request", reqid).Logger()
			return reqid
		},
	}))
	app.Use(middleware.Logging())
	app.Use(helmet.New())
	app.Use(cors.New())

	// Services
	shortLinkService := services.NewShortLinkService(postgres)

	// Routes
	handlers.RegisterShortLinkRoutes(app, shortLinkService)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return utils.NewError(fiber.StatusNotFound, "the requested endpoint does not exist")
	})

	// Start server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	port := viper.GetInt("server.port")

	go func() {
		log.Info().Int("port", port).Msg("api is ready and listening")
		err := app.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			log.Error().Err(err).Msg("error starting api application")
		}
	}()

	<-done

	log.Info().Msg("stopping server...")

	err = app.Shutdown()
	if err != nil {
		log.Error().Err(err).Msg("failed to shut down api server")
	}

}
