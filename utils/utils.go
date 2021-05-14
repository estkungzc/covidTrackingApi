package utils

import (
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClienter interface {
	Do(req *http.Request) (*http.Response, error)
}

type CustomHttpClient struct {
	Client      HttpClienter
	Method      string
	Url         string
	BodyRequest io.Reader
}

func MakeRequest(request CustomHttpClient) ([]byte, error) {
	// Create an http.Request instance
	req, _ := http.NewRequest(request.Method, request.Url, request.BodyRequest)
	// Call the `Do` method, which has a similar interface to the `http.Do` method
	res, err := request.Client.Do(req)
	if err != nil {
		panic(err)
	}

	// Heimdall returns the standard *http.Response object
	body, err := ioutil.ReadAll(res.Body)
	// Close response body
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	return body, nil
}
