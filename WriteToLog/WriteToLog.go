// ready for a log
package writetolog

import (
	"log"
	"os"
)

func WriteToLog() {
	logFile, err := os.OpenFile("../consensus-log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
}
