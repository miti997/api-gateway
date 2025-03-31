package logging

import (
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	l, err := NewDefaultLogger("temp", "test_log", 0)

	if err != nil {
		t.Fatalf("Logger could not be initialized: %s", err)
	}

	defer l.logFile.Close()

	filePath := l.logFile.Name()

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		t.Fatalf("Expected file %s to exist, but it doesn't", filePath)
	} else if err != nil {
		t.Fatalf("Error checking file %s: %v", filePath, err)
	}
}
