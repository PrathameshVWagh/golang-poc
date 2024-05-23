package main

import "fmt"

func main() {

	age := 23
	name := "Prathamesh"
	height := 23.34

	fmt.Println("Name :", name, "Age :", age, "Height :", height)
	fmt.Println("Hello World")

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Height is %.2f\n", height)
	fmt.Printf("Type of name %T\n", name)
}
