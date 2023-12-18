package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// This Function checks for potential insecure design issues on the given website URL.
func InsecureDesign(url string) {

	checkHTTPScheme(url)
	checkSensitiveInformationInURLs(url)
	checkDirectoryListing(url)
	checkCrossOriginResourceSharing(url)
	checkMissingSecurityHeaders(url)

}

// checkHTTPScheme checks if the website uses HTTP instead of HTTPS.
func checkHTTPScheme(url string) {
	if !isHTTPS(url) {
		fmt.Println("The website is using HTTP instead of HTTPS. Insecure design detected.")
	}
}

// checkSensitiveInformationInURLs checks if sensitive information is exposed in URLs.
func checkSensitiveInformationInURLs(url string) {
	sensitiveParams := []string{"password", "apikey", "token", "secret"}
	for _, param := range sensitiveParams {
		if strings.Contains(url, param) {
			fmt.Printf("Sensitive information (%s) is exposed in the URL. Insecure design detected.\n", param)
		}
	}
}

// checkDirectoryListing checks if directory listing is enabled.
func checkDirectoryListing(url string) {
	// Example: Send a GET request to a potential directory and check if it returns a directory listing.
	directoryURL := fmt.Sprintf("%s/some-directory/", url)
	response, err := http.Get(directoryURL)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	if strings.Contains(string(body), "<title>Index of") {
		fmt.Println("Directory listing is enabled. Insecure design detected.")
	}
}

// // hasHeader checks if the specified header is present in the HTTP response.
// func hasHeader(url, header string) bool {
//     response, err := http.Get(url)
//     if err != nil {
//         fmt.Println("Error sending GET request:", err)
//         return false
//     }
//     defer response.Body.Close()

//     // Check if the header is present in the response
//     return response.Header.Get(header) != ""
// }

// checkCrossOriginResourceSharing checks for misconfigurations in CORS headers.
func checkCrossOriginResourceSharing(url string) {
	// Example: Send a cross-origin request and check if the expected CORS headers are present.
	corsTestURL := "https://malicious-attacker.com"
	req, err := http.NewRequest("GET", corsTestURL, nil)
	if err != nil {
		fmt.Println("Error creating CORS test request:", err)
		return
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending CORS test request:", err)
		return
	}
	defer response.Body.Close()

	if response.Header.Get("Access-Control-Allow-Origin") == "*" {
		fmt.Println("Insecure CORS configuration detected.")
	}
}

// checkMissingSecurityHeaders checks if essential security headers are missing.
func checkMissingSecurityHeaders(url string) {
	// Check for the presence of security headers like Content-Security-Policy, X-Content-Type-Options, etc.
	securityHeaders := []string{"Content-Security-Policy", "X-Content-Type-Options", "Strict-Transport-Security"}
	for _, header := range securityHeaders {
		if !hasHeader(url, header) {
			fmt.Printf("Missing security header (%s). Insecure design detected.\n", header)
		}
	}
}

// isHTTPS checks if the URL uses HTTPS.
func isHTTPS(url string) bool {
	return strings.HasPrefix(url, "https://")
}

// hasHeader checks if the specified header is present in the HTTP response.
func hasHeader(url, header string) bool {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return false
	}
	defer response.Body.Close()

	return response.Header.Get(header) != ""
}
