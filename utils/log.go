package utils

import (
	"log"
	"os"
)

const (
	logFilePath   = "/var/log/tellmeall.log"
	logToFileFlag = "LOG_TO_FILE"
	logPrefix     = "tellmeall"
)

var (
	logFile *os.File
	err     error
)

// ConfigureLogging configures logging output based on LOG_TO_FILE variable
func ConfigureLogging() {

	// log to file only if the LOG_TO_FILE var is set
	if MustGetEnv(logToFileFlag, "") == "" {
		log.Printf("%s not set, logging to stdout", logToFileFlag)
		return
	}

	logFile, err = os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error while opening log file: %s - %v", logFilePath, err)
	}

	log.SetPrefix(logPrefix)
	log.SetOutput(logFile)
	log.Printf("Logging to file configured: %s", logFilePath)

}

// FinalizeLogging closes log file if it has been created
func FinalizeLogging() {
	if logFile != nil {
		logFile.Close()
	}
}
