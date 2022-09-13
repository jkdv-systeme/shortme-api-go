package handlers

import (
	"api/internal/http/services"
	"api/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type linksHandler struct {
	shortLinkService services.ShortLinkService
}

func RegisterShortLinkRoutes(app *fiber.App, shortLinkService services.ShortLinkService) {
	handler := &linksHandler{
		shortLinkService: shortLinkService,
	}

	router := app.Group("/links")

	router.Get("/", handler.list)
}

func (h *linksHandler) list(c *fiber.Ctx) error {
	shortLinks, err := h.shortLinkService.GetAllShortLinks()

	if err != nil {
		log.Error().Err(err).Msg("could not get short links from database")
		return utils.NewError(fiber.StatusInternalServerError, "could not get short links")
	}

	// Redundant code, but here for better visibility
	c.Status(fiber.StatusOK)
	return c.JSON(&shortLinks)
}
