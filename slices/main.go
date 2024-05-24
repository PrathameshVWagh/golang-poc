package main

import "fmt"

func main() {

	fmt.Println(".... Slices ....")

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Slices :", numbers)
	fmt.Println("Length :", len(numbers))
	fmt.Println("Capacity :", cap(numbers))

	number1 := append(numbers, 2, 3, 4, 45)
	fmt.Println("Append Slices ", number1)

	num := make([]int, 3, 5)
	num = append(num, 1)
	num = append(num, 3)
	//num = append(num, 4)
	fmt.Println("Golang :", num)
	fmt.Println("Length :", len(num))
	fmt.Println("Capacity :", cap(num))

}
