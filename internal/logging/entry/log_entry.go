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
	SetLatency(time.Time, time.Time)
}

type DefaultLogEntry struct {
	Timestamp  string `json:"timestamp"`
	IP         string `json:"ip"`
	Level      string `json:"level"`
	Request    string `json:"request"`
	PathIn     string `json:"path_in"`
	PathOut    string `json:"path_out"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Latency    int64  `json:"latency"`
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
func (e *DefaultLogEntry) SetLatency(st time.Time, et time.Time) {
	e.Latency = et.Sub(st).Milliseconds()
}
func (e *DefaultLogEntry) SetPath(p string) {
	e.PathIn = p
}
func (e *DefaultLogEntry) SetPathOut(p string) {
	e.PathOut = p
}
