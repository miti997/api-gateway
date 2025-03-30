package routing

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type RouteInterface interface {
	GetPattern() string
}

type Route struct {
	request       string
	in            string
	out           string
	pathParamsIn  map[string]struct{}
	pathParamsOut map[string]struct{}
}

const (
	get    = "GET"
	post   = "POST"
	put    = "PUT"
	patch  = "PATCH"
	delete = "DELETE"
)

func NewRoute(request string, in string, out string) (*Route, error) {
	r := &Route{}

	err := r.setRequest(request)

	if err != nil {
		return nil, err
	}

	err = r.setIn(in)
	if err != nil {
		return nil, err
	}

	err = r.setOut(out)
	if err != nil {
		return nil, err
	}

	r.extractPathParams()

	err = r.comparePathParams()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Route) setRequest(request string) error {
	request = strings.ToUpper(request)

	var validRequests = map[string]struct{}{
		get:    {},
		post:   {},
		put:    {},
		delete: {},
	}

	if _, exists := validRequests[request]; !exists {
		return fmt.Errorf("\"%s\" is not a valid request type", request)
	}
	r.request = request
	return nil
}

func (r *Route) setIn(inUrl string) error {
	if strings.Contains(inUrl, " ") {
		return fmt.Errorf("URL contains spaces, which are not allowed")
	}

	_, err := url.ParseRequestURI(inUrl)
	if err != nil {
		return err
	}
	r.in = inUrl

	return nil
}

func (r *Route) setOut(outUrl string) error {
	if strings.Contains(outUrl, " ") {
		return fmt.Errorf("URL contains spaces, which are not allowed")
	}
	_, err := url.ParseRequestURI(outUrl)
	if err != nil {
		return err
	}
	r.out = outUrl

	return nil
}

func (r *Route) extractPathParams() {
	fn := func(url string) map[string]struct{} {
		re := regexp.MustCompile(`\{([^}]+)\}`)
		matches := re.FindAllStringSubmatch(url, -1)

		paramsMap := make(map[string]struct{})
		for _, match := range matches {
			if len(match) > 1 {
				paramsMap[match[1]] = struct{}{}
			}
		}

		return paramsMap
	}

	r.pathParamsIn = fn(r.in)
	r.pathParamsOut = fn(r.out)
}

func (r *Route) comparePathParams() error {
	lenIn := len(r.pathParamsIn)
	lenOut := len(r.pathParamsOut)

	if lenIn != lenOut {
		return fmt.Errorf("inbound URL has %d path parameters while outbound URL has %d", lenIn, lenOut)
	}

	for param := range r.pathParamsIn {
		if _, exists := r.pathParamsOut[param]; !exists {
			return fmt.Errorf("parameter %s exists in inbound URL but not in outbound URL", param)
		}
	}

	for param := range r.pathParamsOut {
		if _, exists := r.pathParamsIn[param]; !exists {
			return fmt.Errorf("parameter %s exists in outbound URL but not in inbound URL", param)
		}
	}

	return nil
}

func (r *Route) GetPattern() string {
	return fmt.Sprintf("%s %s", r.request, r.in)
}

func (r *Route) HandleRequest(w http.ResponseWriter, req *http.Request) {
	for key := range r.pathParamsOut {
		regex := regexp.MustCompile(`\{` + key + `\}`)
		r.out = regex.ReplaceAllString(r.out, req.PathValue(key))
	}

	outR, err := http.NewRequest(r.request, r.out, req.Body)

	if len(req.URL.Query()) > 0 {
		r.out = r.out + "?" + req.URL.Query().Encode()
	}

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for key, values := range req.Header {
		for _, value := range values {
			outR.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(outR)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Fatalf("Error copying response body: %v", err)
	}
}
