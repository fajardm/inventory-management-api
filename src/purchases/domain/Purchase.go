package domain

import (
	"github.com/fajardm/inventories/src/commons/domain"
)

type Purchase struct {
	domain.Model
	OrderQuantity  int     `gorm:"default:0"`
	NumberReceived int     `gorm:"default:0"`
	Price          float32 `gorm:"default:0"`
	TotalPrice     float32 `gorm:"default:0"`
	Receipt        string  `gorm:"size:255"`
	Note           string  `gorm:"size:255"`
	ProductId      string
}
