package routing

import (
	"fmt"
	"testing"
)

func TestSetRequestSuccess(t *testing.T) {
	request := get
	r := &Route{}

	err := r.setRequest(request)

	if err != nil {
		t.Fatalf("Could not set request: %s", err)
	}

	if request != r.request {
		t.Fatalf("Request was set but improperly")
	}
}

func TestSetRequestFail(t *testing.T) {
	request := "wrong"
	r := &Route{}

	err := r.setRequest(request)

	if err == nil {
		t.Fatalf("Request was set although it shouldn't have been")
	}
}

func TestSetInUrlSuccess(t *testing.T) {
	url := "/some/valid/url"

	r := &Route{}

	err := r.setIn(url)

	if err != nil {
		t.Fatalf("Could not set in url: %s", err)
	}

	if r.in != url {
		t.Fatalf("In url was set but improperly")
	}
}

func TestSetInUrlFail(t *testing.T) {
	url := "3231"

	r := &Route{}

	err := r.setIn(url)

	if err == nil {
		t.Fatalf("In url was set but it shouldn't have been: %s", err)
	}
}

func TestSetOutUrlSuccess(t *testing.T) {
	url := "/some/valid/url"

	r := &Route{}

	err := r.setOut(url)

	if err != nil {
		t.Fatalf("Could not set in url: %s", err)
	}

	if r.out != url {
		t.Fatalf("Out url was set but improperly")
	}
}

func TestSetOutUrlFail(t *testing.T) {
	url := "3231"

	r := &Route{}

	err := r.setOut(url)

	if err == nil {
		t.Fatalf("In url was set but it shouldn't have been: %s", err)
	}
}

func TestExtractPathParams(t *testing.T) {
	extracted1 := "param1"
	extracted2 := "param2"
	notExtracted := "param3"
	url := fmt.Sprintf("/{%s}/{%s}/%s", extracted1, extracted2, notExtracted)

	r := &Route{
		in: url,
	}

	r.extractPathParams()

	if _, exists := r.pathParamsIn[extracted1]; !exists {
		t.Fatalf("%s should have been extracted but wasn't", extracted1)
	}

	if _, exists := r.pathParamsIn[extracted2]; !exists {
		t.Fatalf("%s should have been extracted but wasn't", extracted2)
	}

	if _, exists := r.pathParamsIn[notExtracted]; exists {
		t.Fatalf("%s shouldn't have been extracted but was", notExtracted)
	}
}

func TestComparePathParamsSuccess(t *testing.T) {
	r := &Route{}

	var pathParams = map[string]struct{}{
		"id":   {},
		"name": {},
	}

	r.pathParamsIn = pathParams
	r.pathParamsOut = pathParams

	err := r.comparePathParams()

	if err != nil {
		t.Fatalf("Failed to assert that inbound path params are the same as outbound path params: %s", err)
	}
}

func TestComparePathParamsSail(t *testing.T) {
	r := &Route{}

	var pathParamsIn = map[string]struct{}{
		"id": {},
	}

	r.pathParamsIn = pathParamsIn

	pathParamsOut := map[string]struct{}{
		"fail": {},
	}

	r.pathParamsOut = pathParamsOut

	err := r.comparePathParams()

	if err == nil {
		t.Fatalf("Inbound and outbound parameters should have missmatches but they haven't had any")
	}
}

func TestGetPattern(t *testing.T) {
	request := get
	in := "/test"

	expected := fmt.Sprintf("%s %s", request, in)

	r := &Route{
		request: request,
		in:      in,
	}

	if r.GetPattern() != expected {
		t.Fatalf("Pattern could not be retrieved properly")
	}
}
