package domain

import (
	"github.com/fajardm/inventories/src/commons/domain"
	domain2 "github.com/fajardm/inventories/src/products/domain"
)

type Purchase struct {
	domain.Model
	OrderQuantity  int             `gorm:"default:0"`
	NumberReceived int             `gorm:"default:0"`
	Price          float32         `gorm:"default:0"`
	TotalPrice     float32         `gorm:"default:0"`
	Receipt        string          `gorm:"size:255"`
	Note           string          `gorm:"size:255"`
	Product        domain2.Product `gorm:"index,ForeignKey:ProductId,AssociationForeignKey:id"`
	ProductId      string          `gorm:"index"`
}