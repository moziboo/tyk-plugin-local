package logger

import (
	"log"
	"os"
)

// Logger for simple logging
var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func Info(s string) {
	logger.Println(s)
}
