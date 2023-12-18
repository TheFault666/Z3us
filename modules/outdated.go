package modules

//github.com/otiai10/gosseract refferd for outdated components
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// DetectOutdatedComponents checks for potential vulnerable and outdated components on the given website URL.
func OutdatedComponents(url string) {
	// Example checks for outdated components
	checkTesseractVersion(url)
	// Add more checks as needed
}

// checkTesseractVersion checks the version of an external component (hypothetical example).
func checkTesseractVersion(url string) {
	// Example: Sending a request to the website and analyzing the response for Tesseract version.
	response, err := http.Get(url)
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

	// Regular expression to extract the Tesseract version from the response body
	versionRegex := regexp.MustCompile(`Tesseract Version: (\d+\.\d+\.\d+)`)
	matches := versionRegex.FindStringSubmatch(string(body))

	if len(matches) > 1 {
		tesseractVersion := matches[1]
		fmt.Printf("Detected Tesseract version: %s\n", tesseractVersion)

		// Check if the Tesseract version is outdated (hypothetical condition)
		if isOutdatedTesseractVersion(tesseractVersion) {
			fmt.Println("The Tesseract version is outdated. Vulnerable component detected.")
		} else {
			fmt.Println("The Tesseract version is up-to-date.")
		}
	} else {
		fmt.Println("Could not determine the Tesseract version from the response.")
	}
}

// isOutdatedTesseractVersion checks if the Tesseract version is outdated (hypothetical condition).
func isOutdatedTesseractVersion(version string) bool {
	return version < "4.1.0"
}
