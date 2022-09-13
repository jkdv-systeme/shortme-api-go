package domain

import (
	"time"
)

type IdentityModel struct {
	ID int `json:"id" gorm:"primarykey" swaggertype:"string"`
}

type TemporalModel struct {
	IdentityModel
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
