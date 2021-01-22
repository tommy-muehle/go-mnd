package operation

func example() {
	var y int
	_ = y * 20 // want "Magic number: 20"
}

func example2() {
	var y int
	_ = 10 * y // want "Magic number: 10"
}

func example3() {
	var y int
	_ = 5 * y * 6 // want "Magic number: 5" "Magic number: 6"
}

func example4() {
	const c = 24
	_ = c + (42 * 10) // want "Magic number: 42" "Magic number: 10"
}

func example5() {
	var x int32
	if (42 * x) > 10 { // want "Magic number: 42" "Magic number: 10"
		// ...
	}
}

func example6() {
	var x int32
	if 10 < (x * 42) { // want "Magic number: 42" "Magic number: 10"
		// ...
	}
}

func example7() {
	var x float32
	if 10 < (x * 1.0) { // want "Magic number: 10"
	}
}
