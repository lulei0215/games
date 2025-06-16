package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // ID
	CreatedAt time.Time      //
	UpdatedAt time.Time      //
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` //
}
