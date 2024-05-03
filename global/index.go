package global

import (
	"github.com/sirupsen/logrus"
	"go-admin/config"
	"gorm.io/gorm"
)

var (
	CONFIG *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
