package main

import "fmt"

func main() {
	var i int = 21
	var j bool = true
	var flt float64 = 123.456000
	fmt.Printf("%T", i)
	fmt.Println("\n", j)
	fmt.Println(flt)
	fmt.Printf("%e\n", flt)
}
