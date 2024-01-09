package model

import (
	"github.com/lib/pq"
)

// Person 个人
type Person struct {
	Model
	Name                string          `json:"name" gorm:"not null"`                      // 个人名称
	Email               string          `json:"email" gorm:"not null;uniqueIndex"`         // 邮箱
	Introduction        string          `json:"introduction" gorm:"not null"`              // 个人简介
	AvatarWord          string          `json:"avatarWord" gorm:"size:1"`                  // word of avatar
	AboutMeIntroduction pq.StringArray  `json:"aboutMeIntroduction" gorm:"type:varchar[]"` // 个人简介详情
	AboutMeSkills       pq.StringArray  `json:"aboutMeSkills" gorm:"type:varchar[]"`       // 个人技能
	JobExperiences      []JobExperience `json:"jobExperiences" gorm:"foreignKey:PersonID"` // 工作
	Works               []Work          `json:"works" gorm:"foreignKey:PersonID"`          // 作品
}
