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
	case t.Hour() < 17: // want "Magic number: 17"
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
