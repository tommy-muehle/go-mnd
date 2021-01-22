package _case

import (
	"fmt"
	"time"
)

func example() {
	var x interface{}

	// ...

	switch x {
	case "test":
	case 3: // want "Magic number: 3"
	}
}

func example2() {
	t := time.Now()
	switch {
	case t.Hour() < 12: // want "Magic number: 12"
		fmt.Println("Good morning!")
	case 17 > t.Hour(): // want "Magic number: 17"
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func example3() {
	var x interface{}

	switch x {
	case 1.0:
	case 0.0:
	case 3.0: // want "Magic number: 3.0"
	}
}
