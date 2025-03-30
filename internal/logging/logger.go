package logging

import (
	"github.com/miti997/api-gateway/internal/logging/entry"
	"github.com/miti997/api-gateway/internal/logging/formatter"
)

func Log() {
	jsonFormatter := formatter.JSONFormatter{}

	e := entry.DefaultLogEntry{}

	jsonFormatter.Format(&e)
}
