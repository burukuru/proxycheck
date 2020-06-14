package main

import (
	"fmt"
	"io/ioutil"
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
