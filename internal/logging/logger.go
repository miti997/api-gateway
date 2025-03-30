package logging

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/miti997/api-gateway/internal/logging/entry"
	"github.com/miti997/api-gateway/internal/logging/formatter"
)

type Logger struct {
	maxSize   int
	logFile   *os.File
	fileName  string
	mu        sync.Mutex
	formatter formatter.Formatter
}

func NewLogger(filePath string, fileName string, ms int) (*Logger, error) {
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

	return &Logger{
		logFile:   logFile,
		maxSize:   ms,
		fileName:  fileName,
		formatter: formatter.NewJSONFormatter(),
	}, nil
}

func (l *Logger) Close() error {
	return l.logFile.Close()
}

func (l *Logger) log(level string, ip string, r string, in string, out string, status int, start time.Time, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	e := &entry.DefaultLogEntry{}

	e.SetTimestamp(start)
	e.SetLevel(level)
	e.SetIP(ip)
	e.SetRequest(r)
	e.SetPath(in)
	e.SetPathOut(out)
	e.SetStatusCode(status)
	e.SetLatency(start, time.Now())
	e.SetMessage(message)

	fmt.Println(e.Level)
	log, err := l.formatter.Format(e)

	if err != nil {
		e.SetMessage(fmt.Sprintf("Error formatting log: %v", err))
	}

	_, err = l.logFile.WriteString(log + "\n")
	if err != nil {
		fmt.Printf("Error writing log to file: %v\n", err)
	}
}
