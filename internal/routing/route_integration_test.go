package routing

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/miti997/api-gateway/internal/logging/entry"
)

type TestLogger struct{}

func (l *TestLogger) Log(e entry.LogEntry) {

}

func TestNewRouteSuccess(t *testing.T) {
	request := get
	in := "/testincoming/{id}"
	out := "/testoutgoing/{id}"

	r, err := NewRoute(request, in, out, &TestLogger{})

	if err != nil {
		t.Fatalf("Could not generate new Route: %s", err)
	}

	if r.request != request {
		t.Fatalf("Request was not set properly")
	}

	if r.in != in {
		t.Fatalf("Request was not set properly")
	}

	if r.out != out {
		t.Fatalf("Request was not set properly")
	}
}

func TestNewRouteFailRequest(t *testing.T) {
	request := "wrong"
	in := "/testincoming/{id}"
	out := "/testoutgoing/{id}"

	_, err := NewRoute(request, in, out, &TestLogger{})

	if err == nil {
		t.Fatalf("Route created sucessfully but it shouldn't have been")
	}
}

func TestNewRouteFailIn(t *testing.T) {
	request := get
	in := "/test incoming/{id}"
	out := "/testoutgoing/{id}"

	_, err := NewRoute(request, in, out, &TestLogger{})

	if err == nil {
		t.Fatalf("Route created sucessfully but it shouldn't have been")
	}
}

func TestNewRouteFailOut(t *testing.T) {
	request := "wrong"
	in := "/testincoming/{id}"
	out := "/test outgoing/{id}"

	_, err := NewRoute(request, in, out, &TestLogger{})

	if err == nil {
		t.Fatalf("Route created sucessfully but it shouldn't have been")
	}
}

func TestNewRouteFailPathParams(t *testing.T) {
	request := "wrong"
	in := "/testincoming/{id}"
	out := "/testoutgoing/{id}/{wrong}"

	_, err := NewRoute(request, in, out, &TestLogger{})

	if err == nil {
		t.Fatalf("Route created sucessfully but it shouldn't have been")
	}
}

func TestRouteHandling(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"message": "Success"}`)
	}))
	defer mockServer.Close()

	r := &Route{
		request: get,
		out:     mockServer.URL + "/out/{id}",
		logger:  &TestLogger{},
	}

	req := httptest.NewRequest("GET", "/somepath", nil)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(r.HandleRequest)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK but got %v", rr.Code)
	}

	expectedBody := `{"message": "Success"}`
	if rr.Body.String() != expectedBody+"\n" {
		t.Errorf("Expected response body to be %v but got %v", expectedBody, rr.Body.String())
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type application/json but got %v", rr.Header().Get("Content-Type"))
	}
}
