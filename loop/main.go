package main

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		fmt.Printf("%3d", i)
		i += 1
	}
	fmt.Println()

	for k := 0; k <= 10; k++ {
		fmt.Println(k)
	}

	// Something that I've never seen before:
	if n := 123; n%2 == 0 {
		fmt.Println("N is even")
	} else {
		fmt.Println("N is odd")
	}
}
