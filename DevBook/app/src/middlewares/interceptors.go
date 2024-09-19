package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ValidateResponse(response *http.Response) (*http.Response, error) {
	response = responseLogger(response)
	return response, nil
}

func responseLogger(response *http.Response) *http.Response {
	fmt.Printf("[HttpClient] - \n%s: %s - %s, got %s\n", time.Now().Format("02/01/2006 15:04:05"), response.Request.Method, response.Request.URL, response.Status)
	return response
}
