package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitialSet(file string) {
	fileLog, err := os.OpenFile("../logs/"+file+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error while opening a file: %v\n", err)
	}
	log.SetOutput(fileLog)
	log.SetFormatter(&log.JSONFormatter{})
}
