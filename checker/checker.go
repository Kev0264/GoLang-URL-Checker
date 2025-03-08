package checker

import (
	"net/http"
	"time"
)

// Result holds the outcome of a URL check
type Result struct {
	URL        string
	StatusCode int
	Status	   string
	Error        error
}

// CheckURL sends an HTTP GET request to the given URL and returns the result
func CheckURL(url string) Result {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return Result{URL: url, Error: err}
	}
	defer resp.Body.Close()
	return Result{URL: url, StatusCode: resp.StatusCode, Status: resp.Status}
}

// CheckURLs concurrently checks multiple URLs and returns their results
func CheckURLs(urls []string) []Result {
	results := make([]Result, len(urls))
	resultChan := make(chan Result)

	for i, url := range urls {
		// make sure url is not an empty string
		go func(url string) {
			resultChan <- CheckURL(url)
		}(url)
		results[i] = <- resultChan  //look up <- operator
	}

	return results
}
