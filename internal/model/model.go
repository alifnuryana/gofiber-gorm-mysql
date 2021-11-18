package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primary_key"`
	Title    string
	Subtitle string
	Text     string
}
