package main

func main() {
	// for loop: the only one with semicolons
	for i := 0; i < 5; i++ {
		println(i)
	}

	// create slice and for loop through it
	slice := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}
	for _, v := range slice {
		println(v)
	}

	// "while" loop
	j := 0
	for j < 5 {
		println(j)
		j++
	}

	// infinite loop
	for {
		break
	}

	for true {
		break
	}

	for {
		break
	}

	// if-else
	if 3 < 4 {
		println("3 is less than 4")
	} else {
		println("3 is not less than 4")
	}

	// two statements in one line
	if z := 4; z < 5 {
		println("z is less than 5")
	}

	// switch
	var color string = "red"
	switch color {
	case "green":
		println("color is green")
	case "red":
		println("color is red")
	}

	// switch with default
	switch color {
	case "green":
		println("color is green")
	case "blue":
		println("color is blue")
	default:
		println("color is not green or blue")
	}

	// switch with fallthrough
	// - fallthrough will execute the next case even if it doesn't match
	switch color {
	case "red":
		println("color is red")
		fallthrough
	case "green":
		println("color is green")
	}

	x := 5
	switch {
	case x < 5:
		println("x is less than 5")
		fallthrough
	case x <= 4:
		println("x is less than or equal to 4")
	}
}
