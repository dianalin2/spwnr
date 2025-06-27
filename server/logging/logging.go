package logging

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Log(message string)
}

type FileLogger struct {
	file *os.File
}

func formatLogMessage(message string) string {
	return fmt.Sprintf("[%s]: %s", "INFO", message)
}

func CreateFileLogger(filename string) (*FileLogger, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Failed to open log file: %w", err)
	}
	return &FileLogger{file: file}, nil
}

func (f *FileLogger) Log(message string) {
	logMessage := formatLogMessage(message)
	fmt.Println(logMessage) // Print to console for immediate feedback
	if _, err := f.file.WriteString(logMessage + "\n"); err != nil {
		log.Printf("Failed to write log message: %v", err)
	}
}
