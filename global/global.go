package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Conf *Configs
	Log  *zap.Logger
)
