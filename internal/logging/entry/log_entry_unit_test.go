package entry

import (
	"testing"
	"time"
)

func TestSetTimestamp(t *testing.T) {
	e := &DefaultLogEntry{}

	time := time.Now()

	e.SetTimestamp(time)

	if time.Format("2006-01-02 15:04:05") != e.Timestamp {
		t.Fatalf("Timestamp improperly set")
	}

}
func TestSetIPSuccess(t *testing.T) {
	e := &DefaultLogEntry{}

	ip := "192.168.1.0"

	err := e.SetIP(ip)

	if err != nil {
		t.Fatalf("Failed to set ip: %s", err)
	}

	if e.IP != ip {
		t.Fatalf("Ip not properly set")
	}
}

func TestSetIPFail(t *testing.T) {
	e := &DefaultLogEntry{}

	ip := "Invalid"

	err := e.SetIP(ip)

	if err == nil {
		t.Fatalf("IP was set but it shouldn't have been")
	}
}
func TestSetLevel(t *testing.T) {
	e := &DefaultLogEntry{}

	e.SetLevel(INFO)

	if e.Level != INFO {
		t.Fatalf("Level improperly set")
	}

}
func TestSetRequest(t *testing.T) {
	e := &DefaultLogEntry{}

	r := "GET"

	e.SetRequest(r)

	if e.Request != r {
		t.Fatalf("Request improperly set")
	}

}
func TestSetPath(t *testing.T) {
	e := &DefaultLogEntry{}

	p := "some path"
	e.SetPath(p)

	if e.PathIn != p {
		t.Fatalf("In path set improperly")
	}
}

func TestSetPathPut(t *testing.T) {
	e := &DefaultLogEntry{}

	p := "some path"
	e.SetPathOut(p)

	if e.PathOut != p {
		t.Fatalf("Out path set improperly")
	}
}

func TestSetStatusCode(t *testing.T) {
	e := &DefaultLogEntry{}

	s := 200
	e.SetStatusCode(s)

	if e.StatusCode != s {
		t.Fatalf("Statuc code set improperly")
	}
}
func TestSetMessage(t *testing.T) {
	e := &DefaultLogEntry{}

	m := "message"
	e.SetMessage(m)

	if e.Message != m {
		t.Fatalf("Message set improperly")
	}
}
func TestSetLatency(t *testing.T) {
	e := &DefaultLogEntry{}

	startTime := time.Now()
	e.SetTimestamp(startTime)

	time.Sleep(1 * time.Second)
	e.SetLatency(time.Now())

	latencyDuration, err := time.ParseDuration(e.Latency)
	if err != nil {
		t.Fatalf("Error parsing latency string: %v", err)
	}

	expected := time.Second
	tolerance := 50 * time.Millisecond

	if latencyDuration < expected-tolerance || latencyDuration > expected+tolerance {
		t.Fatalf("Expected Latency around 1s, but got %v", latencyDuration)
	}
}
