package main

import "fmt"

func main() {

	// var declares 1 or more variables of the same type
	var a = "I'm a string"
	fmt.Println(a)

	// can be followed by explicit type
	var num10, num20 int = 10, 20
	fmt.Printf("num10=%d num20=%d\n", num10, num20)

	var num1, num2 = 1, 2
	fmt.Printf("num1=%d num2=%d\n", num1, num2)

	// can't be mixed types though!
	//var num100 bool, num200 int = true, 20
	//fmt.Printf("num100=%s num200=%d\n", num100, num200)

	// This is a shorthand syntax for declaration and initialization
	foo := "Yet another variable"
	fmt.Println(foo)
}
