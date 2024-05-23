package main

import "fmt"

func main() {

	ans := add(10, 23)
	fmt.Println("Addition is :", ans)

	res := sub(23, 12)
	fmt.Println("Subtraction is :", res)
}

func add(a, b int) (result int) {
	return a + b
}

func sub(a int, b int) (result int) {
	result = a - b
	return
}
