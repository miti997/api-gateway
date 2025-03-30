package formatter

import "github.com/miti997/api-gateway/internal/logging/entry"

type Formatter interface {
	Format(le entry.LogEntry) (string, error)
}
