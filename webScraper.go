package main

import (
	"fmt"
	"net/http"
	"io"
	"strings"
	"os"
)

func main() {

	args := os.Args
	//targetUrl := "https://nmap.org"
	targetUrl := args[1]
	fmt.Println("Requesting URL:\t\t", targetUrl)

	response, _ := http.Get(targetUrl)
	fmt.Println("Response Status:\t", response.StatusCode)
	fmt.Println("=========================================")

	body, _ := io.ReadAll(response.Body)
	//fmt.Println(string(body))

	links := string(body)

	lines := strings.Split(links, "\n")

	for _, line := range lines {
		if strings.Contains(line, "href") {
			//fmt.Println("Link: ", line)
			link := strings.Split(line, "\"")
			for _, ln := range link {
				if strings.Contains(ln, "http") {
					fmt.Println("Link: ", ln)
				}
			}
		}
	}
}
