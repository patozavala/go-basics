package main

import "fmt"

func main() {
	// constant declaration
	const pi float64 = 3.14
	const e = 2.71

	fmt.Println("pi:", pi)
	fmt.Println("e:", e)

	// declaration of integer variables
	base := 2
	var height int = 4
	var area int
	area = base * height

	fmt.Println(base, height, area)

	// zero-values
	var zvInt int
	var zvFloat float64
	var zvString string
	var zvBool bool

	fmt.Println(zvInt, zvFloat, zvString, zvBool)

}
