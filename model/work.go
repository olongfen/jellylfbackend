package model

import (
	"github.com/lib/pq"
)

// Work 作品
type Work struct {
	Model
	PersonID uint `json:"personId"`
	// 作品名称
	Name string `json:"name" gorm:"not null;uniqueIndex"`
	// 作品描述
	Descriptions pq.StringArray `json:"descriptions" gorm:"type:varchar[]"`
	// 作品运用技能
	Skills pq.StringArray `json:"skills" gorm:"type:varchar[]"`
}
