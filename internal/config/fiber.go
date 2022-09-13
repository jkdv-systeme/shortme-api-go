package config

import (
	"api/internal/utils"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetFiberConfig() fiber.Config {
	return fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if res, ok := err.(*utils.ServerError); ok {
				ctx.Status(res.Status)
				err := ctx.JSON(&utils.ErrorResponse{
					Status:    res.Status,
					Message:   res.Message,
					RequestID: ctx.Locals("requestid").(string),
				})
				return err
			} else if res, ok := err.(*utils.ValidationError); ok {
				ctx.Status(res.Status)
				err := ctx.JSON(&utils.ErrorResponse{
					Status:    res.Status,
					Message:   res.Message,
					RequestID: ctx.Locals("requestid").(string),
					Fields:    res.Fields,
				})
				return err
			}
			log.Error().Err(err).Msg("Unhandled internal server error")
			ctx.Status(fiber.StatusInternalServerError)
			err = ctx.JSON(&utils.ErrorResponse{
				Status:  fiber.StatusInternalServerError,
				Message: "Something went wrong",
			})
			return err
		},
	}
}
