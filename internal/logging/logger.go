package logging

import (
	"fmt"
	"os"
	"sync"

	"github.com/miti997/api-gateway/internal/logging/entry"
	"github.com/miti997/api-gateway/internal/logging/formatter"
)

type Logger interface {
	Log(entry.LogEntry)
}

type DefaultLogger struct {
	maxSize   int
	logFile   *os.File
	fileName  string
	mu        sync.Mutex
	formatter formatter.Formatter
}

func NewDefaultLogger(filePath string, fileName string, ms int) (*DefaultLogger, error) {
	var logFile *os.File
	var err error

	if filePath == "temp" {
		logFile, err = os.CreateTemp("", "logfile_*.log")
	} else {
		logFile, err = os.OpenFile(filePath+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}

	if err != nil {
		return nil, fmt.Errorf("error opening log file: %v", err)
	}

	if ms < 1 {
		ms = 1
	}

	return &DefaultLogger{
		logFile:   logFile,
		maxSize:   ms,
		fileName:  fileName,
		formatter: formatter.NewJSONFormatter(),
	}, nil
}

func (l *DefaultLogger) Close() error {
	return l.logFile.Close()
}

func (l *DefaultLogger) Log(e entry.LogEntry) {
	l.mu.Lock()
	defer l.mu.Unlock()

	log, err := l.formatter.Format(e)

	if err != nil {
		e.SetMessage(fmt.Sprintf("Error formatting log: %v", err))
	}

	_, err = l.logFile.WriteString(log + "\n")
	if err != nil {
		fmt.Printf("Error writing log to file: %v\n", err)
	}
}
