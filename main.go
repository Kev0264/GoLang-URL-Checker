package main

import (
	"checker/checker"
	"checker/logger"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	urls:=readUrlsFromFile("urls.txt")
	results:=checker.CheckURLs(urls)
	printReport(results)
}

func readUrlsFromFile(filename string) ([]string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed reading file:%v", err)
	}
	urls := strings.Split(string(data), "\n")

	// remove empty strings
	for i := 0; i < len(urls); i++ {
		if urls[i] == "" {
			urls = slices.Delete(urls, i, i+1)
			i--
		}
	}
	return urls
}

func printReport(results []checker.Result) {
	var successCount, failureCount int

	for _, res := range results {
		if res.Error != nil {
			failureCount++
			logger.ErrorLogger.Println(res.Error)
			continue
		} else {
			successCount++
			logger.InfoLogger.Printf("Checked %s: %s\n", res.URL, res.Status)
		}
	}
	fmt.Printf("\nSummary:\nTotal URLs checked:%d\nSuccessful responses:%d\nFailed responses:%d\n",len(results),successCount ,failureCount )
}