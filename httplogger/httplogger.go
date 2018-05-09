package httplogger

import (
	"log"
	"net/http"
	"time"
)

type HTTPLogger interface {
	LogRequest(*http.Request)
	LogResponse(*http.Request, *http.Response, error, time.Duration)
}

type loggedRoundTripper struct {
	rt  http.RoundTripper
	log HTTPLogger
}

func (c *loggedRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	c.log.LogRequest(request)
	startTime := time.Now()
	response, err := c.rt.RoundTrip(request)
	duration := time.Since(startTime)
	c.log.LogResponse(request, response, err, duration)
	return response, err
}

func NewLoggedTransport(rt http.RoundTripper, log HTTPLogger) http.RoundTripper {
	return &loggedRoundTripper{rt: rt, log: log}
}

type DefaultLogger struct {
}

func (dl DefaultLogger) LogRequest(*http.Request) {
}

func (dl DefaultLogger) LogResponse(req *http.Request, res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	if err != nil {
		log.Printf("HTTP Request method=%s host=%s path=%s status=error durationMs=%d error=%q", req.Method, req.Host, req.URL.Path, duration, err.Error())
	} else {
		log.Printf("HTTP Request method=%s host=%s path=%s status=%d durationMs=%d", req.Method, req.Host, req.URL.Path, res.StatusCode, duration)
	}
}

var DefaultLoggedTransport = NewLoggedTransport(http.DefaultTransport, DefaultLogger{})
