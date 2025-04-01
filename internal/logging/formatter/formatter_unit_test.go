package formatter

import (
	"testing"
	"time"
)

type TestEntry struct {
	Request string
	Message string
	Test    string `json:"Test,omitempty"`
}

func (e *TestEntry) SetTimestamp(t time.Time) {}
func (e *TestEntry) SetIP(ip string) error    { return nil }
func (e *TestEntry) SetLevel(l string)        {}
func (e *TestEntry) SetRequest(r string)      {}
func (e *TestEntry) SetStatusCode(s int)      {}
func (e *TestEntry) SetMessage(m string)      {}
func (e *TestEntry) SetLatency(st time.Time)  {}
func (e *TestEntry) SetPath(p string)         {}
func (e *TestEntry) SetPathOut(p string)      {}

func TestFormat(t *testing.T) {
	e := &TestEntry{
		Request: "GET",
		Message: "Success",
	}

	f := JSONFormatter{}
	expected := "{\"Request\":\"GET\",\"Message\":\"Success\"}"
	result, err := f.Format(e)
	t.Log(result)
	if err != nil {
		t.Fatalf("Failed to generate JSON")
	}

	if expected != result {
		t.Fatalf("Json generated but it doesn't match the expected result")
	}
}
