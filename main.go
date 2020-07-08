package main

import "fmt"

//original conditional program
func main() {
	fmt.Println(smaller(5, 7))
}

func smaller(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
