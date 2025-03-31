package logging

import (
	"os"
	"path/filepath"
	"strings"
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

func TestRotateLogFile_Temporary(t *testing.T) {
	logger := &DefaultLogger{
		useTemp: true,
	}

	err := logger.rotateLogFile()
	if err != nil {
		t.Errorf("Expected no error when creating a temp file, got: %v", err)
	}

	if logger.logFile == nil {
		t.Errorf("Log file should be created, but it's nil")
	}

	tempDir := os.TempDir()
	relPath, err := filepath.Rel(tempDir, logger.logFile.Name())
	if err != nil || strings.HasPrefix(relPath, "..") {
		t.Errorf("Temp log file should be inside temp directory, but got: %s", logger.logFile.Name())
	}

	logger.Close()
}

func TestRotateLogFile_Regular(t *testing.T) {
	tmpDir := t.TempDir()
	logger := &DefaultLogger{
		filePath: tmpDir,
		fileName: "testlog",
	}

	err := logger.rotateLogFile()
	if err != nil {
		t.Errorf("Expected no error when creating a regular log file, got: %v", err)
	}

	if logger.logFile == nil {
		t.Errorf("Log file should be created, but it's nil")
	}

	expectedPrefix := filepath.Join(tmpDir, "testlog_")
	if !strings.HasPrefix(logger.logFile.Name(), expectedPrefix) {
		t.Errorf("Log file name should contain timestamped prefix, got: %s", logger.logFile.Name())
	}

	logger.Close()
}

func TestShouldRotate_True(t *testing.T) {
	tmpDir := t.TempDir()
	logger := &DefaultLogger{
		filePath: tmpDir,
		fileName: "testlog",
		maxSize:  10, // Small size for testing
	}

	err := logger.rotateLogFile()
	if err != nil {
		t.Errorf("Failed to create test log file: %v", err)
	}

	_, err = logger.logFile.WriteString("This is a long test entry.")
	if err != nil {
		t.Errorf("Error writing to log file: %v", err)
	}

	if !logger.shouldRotate() {
		t.Errorf("Expected shouldRotate to return true when file size exceeds maxSize")
	}

	logger.Close()
}

func TestShouldRotate_False(t *testing.T) {
	tmpDir := t.TempDir()
	logger := &DefaultLogger{
		filePath: tmpDir,
		fileName: "testlog",
		maxSize:  1000, // 1KB
	}

	err := logger.rotateLogFile()
	if err != nil {
		t.Errorf("Failed to create test log file: %v", err)
	}

	_, err = logger.logFile.WriteString("Short log")
	if err != nil {
		t.Errorf("Error writing to log file: %v", err)
	}

	if logger.shouldRotate() {
		t.Errorf("Expected shouldRotate to return false when file size is below maxSize")
	}

	logger.Close()
}
