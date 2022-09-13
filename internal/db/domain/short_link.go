package domain

// ShortLink godoc
// A short link is a link that redirects to a long link.
//
// @name ShortLink
type ShortLink struct {
	TemporalModel
	ShortCode   string `gorm:"uniqueIndex"`
	OriginalUrl string
}
