package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	baseURL    string
	domainList string
	keywords   string
	outputFile string
	concurrency int
)

func main() {
	flag.StringVar(&baseURL, "u", "", "Single URL to test, e.g., https://example.com")
	flag.StringVar(&domainList, "l", "", "Path to the file containing the list of URLs.")
	flag.StringVar(&keywords, "w", "", "Comma-separated list of keywords to search for in response bodies.")
	flag.StringVar(&outputFile, "o", "", "Output file to save results.")
	flag.IntVar(&concurrency, "c", 10, "Number of concurrent worker threads.")
	flag.Parse()

	if keywords == "" {
		fmt.Println("Usage: ./regexer -u <url> -w <keywords> OR ./regexer -l <file_path> -w <keywords> [-c <concurrency>] [-o <output_file>]")
		os.Exit(1)
	}

	keywordList := strings.Split(keywords, ",")

	if baseURL != "" {
		// Single URL provided via -u flag
		processSingleURL(baseURL, keywordList)
	} else if domainList != "" {
		// URL file provided via -l flag
		urls, err := readURLsFromFile(domainList)
		if err != nil {
			fmt.Println("Error reading URLs from the file:", err)
			os.Exit(1)
		}

		client := &http.Client{Timeout: 5 * time.Second}
		urlsChannel := make(chan string, len(urls))
		resultsChannel := make(chan string, len(urls))
		done := make(chan bool)

		go startWorkerPool(urlsChannel, resultsChannel, concurrency, client, keywordList)
		go processResults(resultsChannel, done, outputFile)

		for _, url := range urls {
			urlsChannel <- url
		}
		close(urlsChannel)

		<-done
	} else {
		fmt.Println("Usage: ./regexer -u <url> -w <keywords> OR ./regexer -l <file_path> -w <keywords> [-c <concurrency>] [-o <output_file>]")
		os.Exit(1)
	}
}

func processSingleURL(url string, keywordList []string) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, readErr := readResponseBodyWithTimeout(resp.Body, 2*time.Second)
	if readErr != nil {
		fmt.Println("Error reading response:", readErr)
		os.Exit(1)
	}

	foundKeywords := checkKeywordsInBody(string(body), keywordList)
	if len(foundKeywords) > 0 {
		fmt.Printf("%s contains: %s\n", url, strings.Join(foundKeywords, ", "))
	}
}

func startWorkerPool(urls <-chan string, results chan<- string, numWorkers int, client *http.Client, keywordList []string) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processURLs(urls, results, client, keywordList)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
}

func processURLs(urls <-chan string, results chan<- string, client *http.Client, keywordList []string) {
	for url := range urls {
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		body, readErr := readResponseBodyWithTimeout(resp.Body, 2*time.Second)
		if readErr != nil {
			continue
		}

		foundKeywords := checkKeywordsInBody(string(body), keywordList)
		if len(foundKeywords) > 0 {
			results <- fmt.Sprintf("%s contains: %s\n", url, strings.Join(foundKeywords, ", "))
		}
	}
}

func processResults(results <-chan string, done chan<- bool, outputFile string) {
	var output *os.File
	var err error

	if outputFile != "" {
		output, err = os.Create(outputFile)
		if err != nil {
			fmt.Println("Error creating output file:", err)
			os.Exit(1)
		}
		defer output.Close()
	}

	for result := range results {
		fmt.Print(result)
		if output != nil {
			_, err := output.WriteString(result)
			if err != nil {
				fmt.Println("Error writing to output file:", err)
			}
		}
	}
	close(done)
}

func readURLsFromFile(filePath string) ([]string, error) {
	var urls []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		urls = append(urls, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func readResponseBodyWithTimeout(body io.Reader, timeout time.Duration) ([]byte, error) {
	done := make(chan struct{})
	var result []byte
	var err error

	go func() {
		defer close(done)
		result, err = io.ReadAll(body)
	}()

	select {
	case <-done:
		return result, err
	case <-time.After(timeout):
		return nil, fmt.Errorf("timeout while reading body")
	}
}

func checkKeywordsInBody(body string, keywordList []string) []string {
	var foundKeywords []string
	for _, keyword := range keywordList {
		if strings.Contains(body, keyword) {
			foundKeywords = append(foundKeywords, keyword)
		}
	}
	return foundKeywords
}
