package main

import "fmt"

func main() {
	var a, b int
	a = 2
	b = 4
	a, b = b, a
	fmt.Println("A:", a)
	fmt.Println("B:", b)
}
