package model

import (
	"github.com/lib/pq"
)

// JobExperience 工作
type JobExperience struct {
	Model
	PersonID uint `json:"personId"`
	// 个人职务
	Position string `gorm:"not null" json:"position"`
	// 任职时间 YYYY-MM
	StartTime string `gorm:"not null" json:"startTime"`
	// 离职时间  YYYY-MM
	EndTime string `gorm:"not null" json:"endTime"`
	// 公司名称
	Company string `gorm:"not null" json:"company"`
	// 公司地址
	CompanyAddress string `json:"companyAddress"`
	// 工作描述
	Descriptions pq.StringArray `gorm:"type:varchar[]" json:"descriptions"`
}
