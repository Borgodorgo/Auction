// ready for a log
package main

import (
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("../consensus-log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
}
