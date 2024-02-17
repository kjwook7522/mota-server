package log

import (
	"log"
	"os"
	"time"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)
var logFile *os.File

func Init() error {

	fileName := "log/go_server_log_" + time.Now().Format("20060102") + ".log"
	fileFlag := os.O_CREATE | os.O_WRONLY | os.O_APPEND
	filePerm := os.FileMode(0666)
	file, err := os.OpenFile(fileName, fileFlag, filePerm)
	if err != nil {
		return err
	}

	Info = log.New(file, "[INFO]", log.LstdFlags)
	Warn = log.New(file, "[WARN]", log.LstdFlags)
	Error = log.New(file, "[ERROR]", log.LstdFlags)
	logFile = file

	return nil
}

func File() *os.File {
	return logFile
}
