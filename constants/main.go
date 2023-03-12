package main

import (
	"fmt"
	"math"
)

func main() {
	const n = 30 + 25i
	fmt.Printf("type=%T value=%v\n", n, n)
	fmt.Println(math.Sin(45))

	// Won't work as n is a constant
	//n = 3 + 1i
	//fmt.Println(n)
}
