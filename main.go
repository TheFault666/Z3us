package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Print("Enter the URL : ")
	var userURL string
	_, _ = fmt.Scanln(&userURL)

	parsedURL, err := url.Parse(userURL)
	if err != nil {
		fmt.Printf("Error parsing the URL : %v\n", err)
		return
	}

	checkSecurityHeaders(parsedURL.String()) // perform basic security  checks
	checkSSL(parsedURL.Host)

}
