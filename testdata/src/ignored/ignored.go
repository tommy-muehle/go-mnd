package ignored

import (
	"net/http"
	"time"
)

func example() interface{} {
	s := struct {
		Amount int
	}{
		Amount: 100,
	}

	return s
}

func example2() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second, // want "Magic number: 5"
	}
}
