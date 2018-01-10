package krakenapi

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func retry(attempts int, sleep time.Duration, callback func() error) (err error) {
	for i := 0; ; i++ {
		err = callback()
		if err == nil {
			return
		}
		if !isRetryable(err) || i >= (attempts-1) {
			break
		}

		time.Sleep(sleep)
		log.Printf("Retry #%d after error: %s\n", i+1, err)
	}
	return err
}

// isRetryable returns true if the error being check can be retried.
func isRetryable(err error) bool {
	// That error happens when the cloudflare page is served.
	if strings.Contains(fmt.Sprintf("%v", err), "Response Content-Type is 'text/html', but should be 'application/json'.") {
		return true
	}
	// That error happens when the service is overloaded.
	if strings.Contains(fmt.Sprintf("%v", err), "EService:Unavailable") {
		return true
	}
	// That error happens when the server response with anything >= 500.
	if strings.Contains(fmt.Sprintf("%v", err), "Response status is an error") {
		return true
	}
	return false
}
