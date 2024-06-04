package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	log.Out = os.Stdout
	log.Level = logrus.InfoLevel
}
