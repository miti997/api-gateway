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

type TestEntry struct{}

func (e *TestEntry) SetTimestamp(t time.Time) {}
func (e *TestEntry) SetIP(ip string) error    { return nil }
func (e *TestEntry) SetLevel(l string)        {}
func (e *TestEntry) SetRequest(r string)      {}
func (e *TestEntry) SetStatusCode(s int)      {}
func (e *TestEntry) SetMessage(m string)      {}
func (e *TestEntry) SetLatency(st time.Time)  {}
func (e *TestEntry) SetPath(p string)         {}
func (e *TestEntry) SetPathOut(p string)      {}

func TestLog(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testlog_*.log")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	logger := &DefaultLogger{
		logFile:   tempFile,
		formatter: &TestFormatter{},
		useTemp:   true,
	}

	entry := &TestEntry{}
	logger.Log(entry)

	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temp log file: %v", err)
	}

	expected := `test` + "\n"

	if string(content) != expected {
		t.Errorf("Log output does not match expected.\nGot: %s\nExpected: %s", content, expected)
	}
}
