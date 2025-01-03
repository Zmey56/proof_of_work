package utils

import "log"

// LogInfo logs informational messages to the console.
func LogInfo(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}
