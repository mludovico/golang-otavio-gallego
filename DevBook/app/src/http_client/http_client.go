package httpclient

import (
	"devbook_app/src/config"
	"devbook_app/src/middlewares"
	"fmt"
	"net/http"
	"strings"
)

type HttpClient struct {
	client       http.Client
	baseURL      string
	authToken    string
	interceptors []func(*http.Request) *http.Request
}

func NewClient(authToken string) *HttpClient {
	return &HttpClient{
		client:    http.Client{},
		baseURL:   config.APIURL,
		authToken: authToken,
	}
}

func (c *HttpClient) Get(route string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, c.baseURL+route, nil)
	if strings.TrimSpace(c.authToken) != "" {
		request.Header.Add("Authorization", "Bearer "+c.authToken)
	}
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return middlewares.ValidateResponse(response)
}

func (c *HttpClient) Post(route string, body string) (*http.Response, error) {
	fmt.Printf("body: %v\n", body)
	request, err := http.NewRequest(http.MethodPost, c.baseURL+route, strings.NewReader(body))
	if strings.TrimSpace(c.authToken) != "" {
		request.Header.Add("Authorization", "Bearer "+c.authToken)
	}
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return middlewares.ValidateResponse(response)
}

func (c *HttpClient) Put(route string, body string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPut, c.baseURL+route, strings.NewReader(body))
	if strings.TrimSpace(c.authToken) != "" {
		request.Header.Add("Authorization", "Bearer "+c.authToken)
	}
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return middlewares.ValidateResponse(response)
}

func (c *HttpClient) Delete(route string, body string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodDelete, c.baseURL+route, strings.NewReader(body))
	if strings.TrimSpace(c.authToken) != "" {
		request.Header.Add("Authorization", "Bearer "+c.authToken)
	}
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return middlewares.ValidateResponse(response)
}
