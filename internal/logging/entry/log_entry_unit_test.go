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
		t.Logf("Timestamp improperly set")
	}

}
func TestSetIPSuccess(t *testing.T) {
	e := &DefaultLogEntry{}

	ip := "192.168.1.0"

	err := e.SetIP(ip)

	if err != nil {
		t.Logf("Failed to set ip: %s", err)
	}

	if e.IP != ip {
		t.Logf("Ip not properly set")
	}
}

func TestSetIPFail(t *testing.T) {
	e := &DefaultLogEntry{}

	ip := "Invalid"

	err := e.SetIP(ip)

	if err == nil {
		t.Logf("IP was set but it shouldn't have been")
	}
}
func TestSetLevel(t *testing.T) {
	e := &DefaultLogEntry{}

	e.SetLevel(INFO)

	if e.Level != INFO {
		t.Logf("Level improperly set")
	}

}
func TestSetRequest(t *testing.T) {
	e := &DefaultLogEntry{}

	r := "GET"

	e.SetRequest(r)

	if e.Request != r {
		t.Logf("Request improperly set")
	}

}
func TestSetPath(t *testing.T) {
	e := &DefaultLogEntry{}

	p := "some path"
	e.SetPath(p)

	if e.PathIn != p {
		t.Logf("In path set improperly")
	}
}

func TestSetPathPut(t *testing.T) {
	e := &DefaultLogEntry{}

	p := "some path"
	e.SetPathOut(p)

	if e.PathOut != p {
		t.Logf("Out path set improperly")
	}
}

func TestSetStatusCode(t *testing.T) {
	e := &DefaultLogEntry{}

	s := 200
	e.SetStatusCode(s)

	if e.StatusCode != s {
		t.Logf("Statuc code set improperly")
	}
}
func TestSetMessage(t *testing.T) {
	e := &DefaultLogEntry{}

	m := "message"
	e.SetMessage(m)

	if e.Message != m {
		t.Logf("Message set improperly")
	}
}
func TestSetLatency(t *testing.T) {
	e := &DefaultLogEntry{}

	startTime := time.Now()

	endTime := startTime.Add(1 * time.Second)

	e.SetLatency(startTime, endTime)

	var expectedLatency = 1000
	if e.Latency != int64(expectedLatency) {
		t.Logf("Expected Latency %v, but got %v", expectedLatency, e.Latency)
	}
}
