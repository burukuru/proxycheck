package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
TODO:
- print out http_proxy and no_prox\
- check urls
  - use timeouts
- print summary of unresponsive urls
- exit 0 if all responded, exit 1 if any did not respond
*/

func main() {
	file := getInputFile(os.Args)
	urls := readFile(file)
	fmt.Println(urls)
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
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	var content []byte
	for {
		buf := make([]byte, 8)
		if _, err := f.Read(buf); err != nil {
			break
		} else {
			content = append(content, buf...)
		}
	}
	urls := strings.Split((string(content)), "\n")
	return urls
}
