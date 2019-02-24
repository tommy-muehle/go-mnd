package _case

func example() {
	var x interface{}

	// ...

	switch x {
	case "test":
	case 3: // want "Magic number: 3"
	}
}
