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

func example4() {
	// ignored via configuration
	a := make([]int, 0, 10)
	a[0] = 1
}
