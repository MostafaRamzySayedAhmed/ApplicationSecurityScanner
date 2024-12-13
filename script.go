package main

import (
	"fmt"
	"net/http"
	"log"
)

func scanHeaders(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error scanning URL %s: %v", url, err)
	}
	defer resp.Body.Close()

	// Check for common security headers
	headers := resp.Header
	if headers.Get("X-Content-Type-Options") == "" {
		fmt.Println("Missing X-Content-Type-Options header (security risk).")
	} else {
		fmt.Println("X-Content-Type-Options header is present.")
	}

	if headers.Get("Strict-Transport-Security") == "" {
		fmt.Println("Missing Strict-Transport-Security header (security risk).")
	} else {
		fmt.Println("Strict-Transport-Security header is present.")
	}

	if headers.Get("X-Frame-Options") == "" {
		fmt.Println("Missing X-Frame-Options header (security risk).")
	} else {
		fmt.Println("X-Frame-Options header is present.")
	}
}

func main() {
	url := "https://example.com"
	fmt.Printf("Scanning %s for security headers...\n", url)
	scanHeaders(url)
}
