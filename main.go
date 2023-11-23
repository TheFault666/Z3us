package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	//Enter a URL
	fmt.Print("Enter a URL: ")
	var userInputURL string
	_, _ = fmt.Scanln(&userInputURL)

	parsedURL, err := url.Parse(userInputURL)
	if err != nil {
		fmt.Printf("Error parsing the URL: %v\n", err)
		return
	}

	//security checks
	checkSecurityHeaders(parsedURL.String())
	checkSSL(parsedURL.Host)
}

func checkSecurityHeaders(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error showing in HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check security headers
	insecureHeaders := []string{
		"X-Content-Type-Options",
		"X-Frame-Options",
		"X-XSS-Protection",
	}

	fmt.Printf("Security Headers:\n")
	for _, header := range insecureHeaders {
		if value := resp.Header.Get(header); value != "" {
			fmt.Printf("%s: %s\n", header, value)
		}
	}
}

func checkSSL(host string) {
	config := &tls.Config{
		InsecureSkipVerify: true, // Skip certificate verification
	}

	conn, err := tls.Dial("tcp", host+":443", config)
	if err != nil {
		fmt.Printf("Error connecting through TLS: %v\n", err)
		return
	}
	defer conn.Close()

	// Check SSL/TLS handshake errors
	if conn.Handshake() != nil {
		fmt.Println("Check certificate and configuration, SSL/TLS Handshake Failed ")
	} else {
		fmt.Println("SSL/TLS Handshake Successful.")
	}
}
