package logging

import (
	"os"
	"testing"
	"time"

	"github.com/miti997/api-gateway/internal/logging/entry"
)

type TestFormatter struct{}

func (f *TestFormatter) Format(e entry.LogEntry) (string, error) {
	return "test", nil
}

func TestLog(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testlog_*.log")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	logger := &Logger{
		logFile:   tempFile,
		formatter: &TestFormatter{},
	}

	startTime := time.Now()
	logger.log(entry.INFO, "127.0.0.1", "GET", "/api/test", "/out", 200, startTime, "Test message")

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temp log file: %v", err)
	}

	expected := `test` + "\n"

	if string(content) != expected {
		t.Errorf("Log output does not match expected.\nGot: %s\nExpected: %s", content, expected)
	}
}
