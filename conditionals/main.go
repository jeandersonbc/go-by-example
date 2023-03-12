package main

import "fmt"

func main() {
	// Something that I've never seen before:
	if n := 123; n%2 == 0 {
		fmt.Println("N is even")
	} else {
		fmt.Println("N is odd")
	}

	// What about a list of initializers?
	//if var a, b, c = 1, 2, 3; if (a + b) > c {
	//    fmt.Println(a, b, c)
	//}

	// And this?
	//if var a = 1; a > 0 {
	//    fmt.Println(a)
	//}

	a, b := 10, 20
	fmt.Println(a, b)

	if c, d := 30, 40; c > d {
		fmt.Println(c)
	} else {
		fmt.Println(d)
	}
}
