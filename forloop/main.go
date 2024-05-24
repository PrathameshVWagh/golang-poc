package main

import "fmt"

func main() {

	fmt.Println("For loop")

	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i)
	}

	counter := 0
	for {
		if counter >= 5 {
			break
		}
		fmt.Println("Counter", counter)
		counter++

	}

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for index, value := range number {
		fmt.Printf("index %d value %d \n", index, value)
	}

	data := "Hello World!!"
	for index, char := range data {
		fmt.Printf("INDEX IS %d AND VALUE IS %c\n", index, char)
	}
}
