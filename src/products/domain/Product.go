package domain

import (
	"github.com/fajardm/inventories/src/commons/domain"
)

type Product struct {
	domain.Model
	SKU   string `gorm:"size:255"`
	Name  string `gorm:"size:255"`
	Stock int    `gorm:"default:0"`
}
