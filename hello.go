package main

import "fmt"

const x = 212.0

func main() {
	fmt.Println("hello,world!")
	var y = x
	var z = (y - 32) * 5 / 9
	fmt.Printf("boilint point = %g or %g\n", z, y)

}
