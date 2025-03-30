package formatter

import (
	"encoding/json"

	"github.com/miti997/api-gateway/internal/logging/entry"
)

type JSONFormatter struct {
}

func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

func (f *JSONFormatter) Format(le entry.LogEntry) (string, error) {
	logEntryJSON, err := json.Marshal(le)
	if err != nil {
		return "", err
	}

	return string(logEntryJSON), nil
}
