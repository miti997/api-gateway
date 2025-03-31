package routing

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/miti997/api-gateway/internal/logging"
	"github.com/miti997/api-gateway/internal/logging/entry"
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
	logger        logging.Logger
}

const (
	get    = "GET"
	post   = "POST"
	put    = "PUT"
	patch  = "PATCH"
	delete = "DELETE"
)

func NewRoute(request string, in string, out string, l logging.Logger) (*Route, error) {
	r := &Route{}

	err := r.setRequest(request)
	r.logger = l

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
	le := entry.NewDefaultLogEntry()

	start := time.Now()

	le.SetTimestamp(start)
	le.SetRequest(r.request)
	le.SetPath(r.in)

	out := r.out
	le.SetIP(req.RemoteAddr)

	for key := range r.pathParamsOut {
		regex := regexp.MustCompile(`\{` + key + `\}`)
		out = regex.ReplaceAllString(out, req.PathValue(key))
	}

	outR, err := http.NewRequest(r.request, out, req.Body)

	if len(req.URL.Query()) > 0 {
		out = r.out + "?" + req.URL.Query().Encode()
	}

	le.SetPathOut(out)

	logFatal := func(message string) {
		le.SetLevel(entry.FATAL)
		le.SetMessage(message)
		le.SetStatusCode(500)
		le.SetLatency(start, time.Now())

		r.logger.Log(le)
	}

	if err != nil {
		logFatal(fmt.Sprintf("Error creating request: %v", err))

		return
	}

	for key, values := range req.Header {
		for _, value := range values {
			outR.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(outR)

	if err != nil {
		logFatal(fmt.Sprintf("Error making request: %s", err))

		return
	}

	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logFatal(fmt.Sprintf("Error reading response body: %v", err))

		return
	}

	bodyString := string(bodyBytes)

	w.WriteHeader(resp.StatusCode)
	le.SetStatusCode(resp.StatusCode)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		le.SetLevel(entry.INFO)
		le.SetMessage("Success")
	} else {
		le.SetLevel(entry.ERROR)
		le.SetMessage(bodyString)
	}

	_, err = w.Write(bodyBytes)
	if err != nil {
		logFatal(fmt.Sprintf("Error writing response body: %v", err))

		return
	}

	r.logger.Log(le)
}
