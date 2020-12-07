package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Log() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)

	return log
}
