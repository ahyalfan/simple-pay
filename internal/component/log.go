package component

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

var Log = initializLogger()

func initializLogger() *logrus.Logger {
	logger := logrus.New()

	file, err := os.OpenFile("logrus.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetOutput(file)
	logger.SetFormatter(&ecslogrus.Formatter{})
	logger.SetLevel(logrus.TraceLevel)
	return logger
}
