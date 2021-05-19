package unirest

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Request struct
type Request struct {
	BaseURL     string
	Endpoint    string
	QueryString string
	Headers     map[string][]string
}

// Response struct
type Response struct {
	RawBody  []byte
	Body     Body
	Headers  map[string][]string
	Status   int
	Protocol string
}

// Body struct that goes in Response
type Body struct {
	Bytes  []byte
	String string
}

// Get : func
func Get(req Request) (Response, error) {

	// Build the URL
	var url string
	url += req.BaseURL
	url += req.Endpoint
	url += req.QueryString

	// Create an HTTP client
	c := &http.Client{Timeout: 30 * time.Second}

	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add any defined headers
	if req.Headers != nil {
		r.Header = http.Header(req.Headers)
	}

	// Add User-Agent if none is given
	if r.Header["User-Agent"] == nil {
		r.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.94 Safari/537.36")
	}

	// Send the request
	res, err := c.Do(r)

	// Check for error
	if err != nil {
		return Response{}, err
	}

	// Make sure to close after reading
	defer res.Body.Close()

	// Limit response body to 1mb
	// lr := &io.LimitedReader{res.Body, 1000000}

	// Read all the response body
	rb, err := ioutil.ReadAll(res.Body)

	// Check for error
	if err != nil {
		return Response{}, err
	}

	// Build the output
	responseOutput := Response{
		RawBody: rb,
		Body: Body{
			Bytes:  rb,
			String: string(rb),
		},
		Headers:  res.Header,
		Status:   res.StatusCode,
		Protocol: res.Proto,
	}

	// Send it along
	return responseOutput, nil

}
