package main

import "fmt"

func main() {

	var name [3]string
	name[0] = "Hello"
	name[1] = "Hi"
	name[2] = "Hello"
	fmt.Println("Name", name)

	var numbers [5]string
	fmt.Printf("Numbers : %q \n", numbers)
}
