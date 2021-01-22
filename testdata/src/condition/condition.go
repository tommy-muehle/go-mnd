package condition

func example() {
	var x int
	var y string

	// ...

	if x > 7 { // want "Magic number: 7"
	}

	if 8 > x { // want "Magic number: 8"
	}

	if "test" == y {
	}
}

func example2() {
	var x float32

	if x > 1.0 {
	}

	if x < 0.0 {
	}
}
