package main

import "fmt"

func main() {
	a := less(5, 6)
	fmt.Println(a)
}

func less(a int, b int) bool {
	return a < b
}
