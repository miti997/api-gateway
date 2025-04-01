package entry

import (
	"fmt"
	"net"
	"time"
)

type LogEntry interface {
	SetTimestamp(time.Time)
	SetIP(string) error
	SetLevel(string)
	SetRequest(string)
	SetPath(string)
	SetStatusCode(int)
	SetMessage(string)
	SetLatency(time.Time)
}

type DefaultLogEntry struct {
	Timestamp  string `json:"timestamp,omitempty"`
	IP         string `json:"ip,omitempty"`
	Level      string `json:"level,omitempty"`
	Request    string `json:"request,omitempty"`
	PathIn     string `json:"path_in,omitempty"`
	PathOut    string `json:"path_out,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
	Latency    string `json:"latency,omitempty"`
}

func NewDefaultLogEntry() *DefaultLogEntry {
	return &DefaultLogEntry{}
}

func (e *DefaultLogEntry) SetTimestamp(t time.Time) {
	e.Timestamp = t.Format("2006-01-02 15:04:05")
}

func (e *DefaultLogEntry) SetIP(ip string) error {
	if net.ParseIP(ip) != nil {
		e.IP = ip
		return nil
	}

	return fmt.Errorf("%s is not a valid IP", ip)
}
func (e *DefaultLogEntry) SetLevel(l string) {
	e.Level = l
}
func (e *DefaultLogEntry) SetRequest(r string) {
	e.Request = r
}
func (e *DefaultLogEntry) SetStatusCode(s int) {
	e.StatusCode = s
}
func (e *DefaultLogEntry) SetMessage(m string) {
	e.Message = m
}
func (e *DefaultLogEntry) SetLatency(st time.Time) {
	duration := time.Since(st)
	e.Latency = duration.String()
}
func (e *DefaultLogEntry) SetPath(p string) {
	e.PathIn = p
}
func (e *DefaultLogEntry) SetPathOut(p string) {
	e.PathOut = p
}
