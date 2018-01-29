package domain

import (
	"github.com/fajardm/inventories/src/commons/domain"
)

type Sale struct {
	domain.Model
	NumberShipped int     `gorm:"default:0"`
	Price         float32 `gorm:"default:0"`
	TotalPrice    float32 `gorm:"default:0"`
	Note          string  `gorm:"size:255"`
	ProductId     string
}
