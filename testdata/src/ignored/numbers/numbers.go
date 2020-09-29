package numbers

import (
	"log"
	"net/http"
	"time"
)

// ignored via configuration
var numberWithDigitSeparator int = 1_000

func example() interface{} {
	s := struct {
		Amount int
	}{
		Amount: numberWithDigitSeparator,
	}

	return s
}

func example2() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second, // want "Magic number: 5"
	}
}

func example3() {
	// ignored via configuration
	if int32(1234_567_890)-int32(1_234_567_890) == 0 {
		log.Println(3.1415_9265) // want "Magic number: 3.1415_9265"
	}
}

func example4() {
	// ignored via configuration
	_ = float32(3.1415_9264)
}
