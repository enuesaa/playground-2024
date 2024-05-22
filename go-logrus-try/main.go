package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	logger := logrus.New()
	logger.Log(logrus.InfoLevel, "hello")

	time.Sleep(2 * time.Second)

	logrus.WithFields(logrus.Fields{
		"a": "b",
	}).Info("c")
}
