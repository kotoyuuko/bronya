package logger

import (
	"io"
	"log"
	"os"
)

var (
	// Info logger
	Info *log.Logger
	// Warning logger
	Warning *log.Logger
	// Error logger
	Error *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	Info = log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, "[WARN]", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "[ERR]", log.Ldate|log.Ltime|log.Lshortfile)
}
