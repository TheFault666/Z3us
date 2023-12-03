package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Print("Enter a website URL: ")
	var userInputURL string
	_, _ = fmt.Scanln(&userInputURL)

	resp, err := http.Get(userInputURL)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check for security headers
	checkSecurityHeaders(resp.Header)
}

func checkSecurityHeaders(headers http.Header) {
	// List of common security headers
	securityHeaders := []string{
		"Strict-Transport-Security",
		"Content-Security-Policy",
		"X-Content-Type-Options",
		"X-Frame-Options",
		"X-XSS-Protection",
	}

	fmt.Println("Security Headers:")
	for _, header := range securityHeaders {
		value := headers.Get(header)
		if value != "" {
			fmt.Printf("%s: %s\n", header, value)
		} else {
			fmt.Printf("%s: Not set\n", header)
		}
	}
}
