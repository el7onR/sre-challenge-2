package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Env struct {
	DB  *gorm.DB
	Log *logrus.Logger
}
