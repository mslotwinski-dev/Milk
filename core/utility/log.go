package utility

import (
	"fmt"
	"log"
)

// Color constants
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
)

func Info(format string, v ...interface{}) {
	log.Println(ColorCyan + "[INFO] " + ColorReset + fmt.Sprintf(format, v...))
}

func Warn(format string, v ...interface{}) {
	log.Println(ColorYellow + "[WARN] " + ColorReset + fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	log.Println(ColorRed + "[ERROR] " + ColorReset + fmt.Sprintf(format, v...))
}
