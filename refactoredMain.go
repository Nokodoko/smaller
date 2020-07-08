package main

import "fmt"

//original conditional program
func main() {
	fmt.Println(evenSmaller(10, 10))
	fmt.Println(evenSmaller(7, 10))
	fmt.Println(evenSmaller(10, 7))
}

func evenSmaller(a, b int) int {
	switch {
	case a < b:
		return a
	case a > b:
		return b
	}
	return 0
}
