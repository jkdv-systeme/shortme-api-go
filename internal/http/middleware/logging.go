package middleware

import (
	"api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"time"
)

func Logging() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		log.Info().Str("method", c.Method()).Str("path", c.Path()).Msg("Received request")
		// Go to next middleware:
		next := c.Next()
		duration := time.Now().Sub(start)
		if next != nil {
			if err, ok := next.(*utils.ServerError); ok {
				log.Info().Interface("request", c.Locals("requestid")).Int("status", err.Status).Str("method", c.Method()).Str("path", c.Path()).Str("duration", duration.String()).Msg("Handled request")
				return err
			}
		} else {
			log.Info().Interface("request", c.Locals("requestid")).Int("status", c.Response().StatusCode()).Str("method", c.Method()).Str("path", c.Path()).Str("duration", duration.String()).Msg("Handled request")
		}
		return next
	}
}
