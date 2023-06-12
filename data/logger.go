package data

import (
	"log"
	"os"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

// CreateLoggers creates all loggers
func CreateLoggers() {
	// Create loggers
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// GetWarningLogger returns the warning logger
func GetWarningLogger() *log.Logger {
	return warningLogger
}

// GetInfoLogger returns the info logger
func GetInfoLogger() *log.Logger {
	return infoLogger
}

// GetErrorLogger returns the error logger
func GetErrorLogger() *log.Logger {
	return errorLogger
}

// CloseLoggers closes all loggers by freeing their memory
func CloseLoggers() {
	warningLogger = nil
	infoLogger = nil
	errorLogger = nil
}
