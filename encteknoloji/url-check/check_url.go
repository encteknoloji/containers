package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"crypto/tls"
)

const (
	defaultInterval    = 3     // Default check interval in seconds
	defaultStatusCode  = http.StatusOK // Default expected status code
	defaultSkipVerify   = false // Default to disabling insecure skips
)

func main() {
	fmt.Println(`
    _______   ________   ______     __               __        _ _ 
   / ____/ | / / ____/  /_  __/__  / /______  ____  / /___    (_|_)
  / __/ /  |/ / /        / / / _ \/ //_/ __ \/ __ \/ / __ \  / / / 
 / /___/ /|  / /___     / / /  __/ ,< / / / / /_/ / / /_/ / / / /  
/_____/_/ |_/\____/    /_/  \___/_/|_/_/ /_/\____/_/\____/_/ /_/   
                                                        /___/ 

    https://www.encteknoloji.com
`)

	url := os.Getenv("CHECK_URL") // Get the URL from environment variable

	if url == "" {
		fmt.Println("Error: CHECK_URL environment variable not set.")
		return
	}

	intervalStr := os.Getenv("CHECK_INTERVAL") // Get the interval from environment variable

	var interval int
	var err error
	if intervalStr == "" { // Use default if not set
		interval = defaultInterval
	} else {
		interval, err = strconv.Atoi(intervalStr) // Convert string to integer
	}

	if err != nil || interval <= 0 {
		fmt.Println("Error: CHECK_INTERVAL invalid (must be a positive integer).")
		return
	}

	intervalDuration := time.Duration(interval) * time.Second

	statusCodeStr := os.Getenv("CHECK_STATUS_CODE") // Get the expected status code from environment variable

	var expectedStatusCode int
	if statusCodeStr == "" { // Use default if not set
		expectedStatusCode = defaultStatusCode
	} else {
		var err error
		expectedStatusCode, err = strconv.Atoi(statusCodeStr) // Convert string to integer
		if err != nil || expectedStatusCode <= 0 {
			fmt.Println("Error: CHECK_STATUS_CODE invalid (must be a positive integer).")
			return
		}
	}

	insecureSkipVerifyStr := os.Getenv("INSECURE_SKIP_VERIFY")
	insecureSkipVerify := defaultSkipVerify // Set default value

	// Convert string to bool (handle potential errors gracefully)
	if insecureSkipVerifyStr != "" {
		var err error
		insecureSkipVerify, err = strconv.ParseBool(insecureSkipVerifyStr)
		if err != nil {
			fmt.Printf("Error converting INSECURE_SKIP_VERIFY value: %v\n", err)
			// Handle error (e.g., exit with appropriate code)
		}
	}

	// Print insecureSkipVerify value in a user-friendly format
	var skipVerifyStatus string
	if insecureSkipVerify {
		skipVerifyStatus = "enabled (WARNING: use with caution!)"
	} else {
		skipVerifyStatus = "disabled"
	}


	fmt.Printf("Using URL: %s\n", url)
	fmt.Printf("Using interval: %d seconds\n", interval)
	fmt.Printf("Using expected status code: %d\n", expectedStatusCode)
	fmt.Printf("using skip certificate verification: %s\n", skipVerifyStatus)
	fmt.Println("")

	// Create the HTTP client
	client := &http.Client{
		Transport: &http.Transport{
		  TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
		},
	}

	for {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("Error getting URL: %v\n", err)
		} else {
			defer resp.Body.Close()

			if resp.StatusCode == expectedStatusCode {
				fmt.Printf("Received expected status code: %d\n", expectedStatusCode)
				break
			} else {
				fmt.Printf("Received status code: %d\n", resp.StatusCode)
			}
		}

		time.Sleep(intervalDuration)
	}
}
