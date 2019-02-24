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
