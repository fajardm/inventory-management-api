package domain

import (
	"time"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (user *Model) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", uuid.New().String())
	return err
}
