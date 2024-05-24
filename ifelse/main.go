package main

import "fmt"

func main() {

	val := 8
	if val > 9 {
		fmt.Println(">9")
	} else if val < 9 {
		fmt.Println("<9")
	} else {
		fmt.Println("=9")
	}

	if val > 9 && (val < 9 || val == 9) {
		fmt.Println("not 9")
	} else {
		fmt.Println("(any)")
	}
}
