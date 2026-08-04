package logger

import (
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger

	logChan chan logMessage
)

type logMessage struct {
	level string
	text  string
}

func Init(infoHandler io.Writer, errorHandler io.Writer) {

	Info = log.New(infoHandler, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandler, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	logChan = make(chan logMessage, 100)
}

func InitDefault() {
	Init(os.Stdout, os.Stderr)
}

func Start() {

	go func() {

		for msg := range logChan {

			switch msg.level {

			case "info":
				Info.Println(msg.text)

			case "error":
				Error.Println(msg.text)
			}
		}
	}()
}

func Stop() {
	close(logChan)
}

func InfoLog(msg string) {
	logChan <- logMessage{
		level: "info",
		text:  msg,
	}
}

func ErrorLog(msg string) {
	logChan <- logMessage{
		level: "error",
		text:  msg,
	}
}