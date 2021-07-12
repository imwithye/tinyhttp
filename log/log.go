package log

import (
	"fmt"
	"os"
	"time"
)

var TraceMode = false

const timeFormat = "2006/01/02 - 15:04:05"

func Info(message string) {
	fmt.Fprintf(os.Stdout, "[INFO] %s | %s\n", time.Now().Format(timeFormat), message)
}

func Warn(message string) {
	fmt.Fprintf(os.Stdout, "[WARN] %s | %s\n", time.Now().Format(timeFormat), message)
}

func Error(message string) {
	fmt.Fprintf(os.Stderr, "[ERRO] %s | %s\n", time.Now().Format(timeFormat), message)
}

func Fatal(message string) {
	fmt.Fprintf(os.Stderr, "[FATA] %s | %s\n", time.Now().Format(timeFormat), message)
	os.Exit(1)
}

func Trace(message string) {
	if TraceMode {
		fmt.Fprintf(os.Stdout, "[TRAC] %s | %s\n", time.Now().Format(timeFormat), message)
	}
}
