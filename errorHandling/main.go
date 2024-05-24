package main

import "fmt"

func main() {

	ans, err := divide(13, 0)
	if err != nil {
		fmt.Println("Error Handling!")
	}
	fmt.Println("Answer is", ans)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil

}
