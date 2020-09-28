package functions

import (
	"math"
	"net/http"
)

func example1() {
	http.StatusText(200) // want "Magic number: 200"
}

func example2() {
	// ignored via configuration
	math.Abs(1.0)
}

func example3() {
	// ignored via configuration
	math.Acos(1.5)
}
