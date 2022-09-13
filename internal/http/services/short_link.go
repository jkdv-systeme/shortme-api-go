package services

import (
	"api/internal/db/domain"
	"gorm.io/gorm"
)

type ShortLinkService interface {
	GetAllShortLinks() ([]domain.ShortLink, error)
}

type shortLinkService struct {
	db *gorm.DB
}

func NewShortLinkService(db *gorm.DB) ShortLinkService {
	return &shortLinkService{
		db: db,
	}
}

// GetAllShortLinks godoc
// Gets all shortened links from the database and returns them in a slice.
// May return an error when the query fails.
func (s *shortLinkService) GetAllShortLinks() ([]domain.ShortLink, error) {
	var shortLinks []domain.ShortLink
	err := s.db.Find(&shortLinks).Error
	return shortLinks, err
}
