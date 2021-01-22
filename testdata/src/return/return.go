package _return

func example() int {
	return 3 // want "Magic number: 3"
}

func example2() string {
	return "3"
}

func example3() float64 {
	return 2.0 // want "Magic number: 2.0"
}

func example4(x int32) int32 {
	return x + 42 // want "Magic number: 42"
}

func example5(x int32) int32 {
	return 42 + x // want "Magic number: 42"
}

func example6(x int32) int32 {
	return x + (42 * 10) // want "Magic number: 42" "Magic number: 10"
}

func example7(x int32) int32 {
	return (42 * x) + 10 // want "Magic number: 42" "Magic number: 10"
}

func example8() float32 {
	return 0.0
}

func example9() float32 {
	return 1.0
}

func example10() float32 {
	return 3.0 // want "Magic number: 3.0"
}
