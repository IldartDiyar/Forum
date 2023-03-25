package Log

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

type Logger struct {
	file *os.File
}

func NewFileLogger() (*Logger, error) {
	file, err := os.OpenFile("logger.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	return &Logger{file}, nil
}

func (l *Logger) Debug(message string) {
	msg := fmt.Sprintf("[DEBUG] %s\n", message)
	output(msg)
}

func (l *Logger) Info(message string) {
	_, file, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[INFO] (file: %s, line: %d): %s\n", file, line, message)
	output(msg)
}

func (l *Logger) Error(message string) {
	_, file, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("[ERROR] (file: %s, line: %d): %s\n", file, line, message)
	output(msg)
}

func (l *Logger) Fatal(message string) {
	l.Error(message)
	os.Exit(1)
}

func output(msg string) {
	log.Print(msg)
}
