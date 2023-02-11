package message

import (
	"io"
	"log"
	"os"
	"strings"

	"github.com/vitorsiqueirarecife/bff/model"
)

type Store interface {
	Save(message model.Message, logMessage, topic string) error
}

type storeImpl struct{}

func NewApp() Store {
	return &storeImpl{}
}

func (a *storeImpl) Save(message model.Message, logMessage, topic string) error {

	topicUpper := strings.ToUpper(topic)
	pathLogs := "./logs/" + topicUpper + "-LOGS"
	pathNew := "./logs/" + topicUpper + "-NEW"

	new, err := os.OpenFile(pathNew, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(new)
	log.Println(logMessage)

	logs, err := os.OpenFile(pathLogs, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	_, err = io.Copy(new, logs)
	if err != nil {
		return err
	}

	new.Close()
	logs.Close()
	os.Remove(pathLogs)
	os.Rename(pathNew, pathLogs)
	return nil
}
