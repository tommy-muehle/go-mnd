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
