package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type urlStatus struct {
	URL        string `json: url`
	StatusCode int    `json: statusCode`
	GetError   error  `json: getError`
}

func main() {
	file := getInputFile(os.Args)
	urls := readFile(file)
	results := checkUrls(urls)
	error_count := printResults(results)
	if error_count > 0 {
		os.Exit(1)
	}
}

func printResults(results []urlStatus) int {
	_, error_count := Summary(results)
	for _, u := range results {
		if u.GetError != nil {
			u.String()
		}
	}
	return error_count
}

func (u *urlStatus) String() {
	if u.GetError != nil {
		fmt.Printf("Failed to get %v (%v): %v\n", u.URL, u.StatusCode, u.GetError)
	} else {
		fmt.Printf("Successful response from %v (%v).\n", u.URL, u.StatusCode)
	}
}

func Summary(r []urlStatus) (int, int) {
	error_count := 0
	success_count := 0
	for _, u := range r {
		if u.GetError != nil {
			error_count = error_count + 1
		} else {
			success_count = success_count + 1
		}
	}
	fmt.Printf("\nSummary:\n")
	fmt.Printf("Success: %v\n", success_count)
	fmt.Printf("Error:   %v\n", error_count)

	return success_count, error_count
}

// Check arguments and return input file
func getInputFile(args []string) string {
	if len(args) != 2 {
		log.Fatal("Please specify file containing list of URLs as argument.")
	}
	inputFile := os.Args[1]
	log.Println("Using URLs from " + inputFile + ".")
	return inputFile
}

// Read file and return URLs within
func readFile(file string) []string {
	var content []byte
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	urls := strings.Split((string(content)), "\n")

	// In POSIX files, splitting by '\n' creates and empty item
	if urls[len(urls)-1] == "" {
		return (urls[:len(urls)-1])
	}
	return urls
}

func checkUrls(urls []string) []urlStatus {
	pf := http.ProxyFromEnvironment

	// Print proxy URL for debugging
	req, _ := http.NewRequest("GET", urls[0], nil)
	p, err := pf(req)
	if err != nil {
		log.Fatal("Could not obtain proxy URL from env", err)
	}

	if p != nil {
		log.Println("Proxy used:", p)
	} else {
		log.Println("WARNING: Proxy URL not found.")
	}

	// Client using proxy from env
	transport := http.Transport{Proxy: pf}
	client := http.Client{
		Timeout:   5 * time.Second,
		Transport: &transport,
	}

	// Run URL checks
	msgs := make(chan urlStatus)
	for _, url := range urls {
		go checkUrl(url, msgs, &client)
	}

	var results []urlStatus
	for i := 0; i < len(urls); i++ {
		results = append(results, <-msgs)
	}

	return results
}

func checkUrl(url string, msgs chan urlStatus, client *http.Client) {
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(url + " F")
		msgs <- urlStatus{url, 0, err}

		return
	}
	fmt.Printf("%v S (%v)\n", url, res.StatusCode)
	msgs <- urlStatus{url, res.StatusCode, nil}
}
