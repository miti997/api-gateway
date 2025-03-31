package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/miti997/api-gateway/internal/logging/entry"
	"github.com/miti997/api-gateway/internal/logging/formatter"
)

type Logger interface {
	Log(entry.LogEntry)
}

type DefaultLogger struct {
	maxSize   int64
	logFile   *os.File
	filePath  string
	fileName  string
	useTemp   bool
	mu        sync.Mutex
	formatter formatter.Formatter
}

func NewDefaultLogger(filePath string, fileName string, maxSizeMB int) (*DefaultLogger, error) {
	if maxSizeMB < 1 {
		maxSizeMB = 1
	}

	useTemp := filePath == "temp"

	logger := &DefaultLogger{
		filePath:  filePath,
		fileName:  fileName,
		useTemp:   useTemp,
		maxSize:   int64(maxSizeMB * 1024 * 1024), // Convert MB to bytes
		formatter: formatter.NewJSONFormatter(),
	}

	err := logger.rotateLogFile()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func (l *DefaultLogger) rotateLogFile() error {
	if l.logFile != nil {
		l.logFile.Close()
	}

	var logFilePath string
	var err error

	if l.useTemp {
		l.logFile, err = os.CreateTemp("", "logfile_*.log")
		if err != nil {
			return fmt.Errorf("error creating temporary log file: %v", err)
		}
	} else {
		if err := os.MkdirAll(l.filePath, os.ModePerm); err != nil {
			return fmt.Errorf("error creating log directory: %v", err)
		}

		logFilePath = filepath.Join(l.filePath, fmt.Sprintf("%s_%s.log", l.fileName, time.Now().Format("20060102_150405")))
		l.logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("error opening log file: %v", err)
		}
	}

	return nil
}

func (l *DefaultLogger) shouldRotate() bool {
	if l.useTemp || l.logFile == nil {
		return false
	}

	info, err := l.logFile.Stat()
	if err != nil {
		return false
	}

	return info.Size() >= l.maxSize
}

func (l *DefaultLogger) Log(e entry.LogEntry) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.shouldRotate() {
		l.rotateLogFile()
	}

	log, err := l.formatter.Format(e)
	if err != nil {
		e.SetMessage(fmt.Sprintf("Error formatting log: %v", err))
	}

	_, err = l.logFile.WriteString(log + "\n")
	if err != nil {
		fmt.Printf("Error writing log to file: %v\n", err)
	}
}

func (l *DefaultLogger) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}
