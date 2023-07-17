package main

import (
	"fmt"
)

func main() {
	var a bool = true // Boolean

	b := 100 // Integer

	var (
		c float32 = 3.14  // Floating point number
		d string  = "Hi!" // String
	)

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
	fmt.Println("map:", m)
}
