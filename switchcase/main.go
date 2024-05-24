package main

import "fmt"

func main() {

	fmt.Println("Hello I am switch")
	day := "susn"

	switch day {
	case "monday", "tuesday", "wednesday":
		fmt.Println("It's is okay")
	case "thrusday", "friday":
		fmt.Println("Great")
	case "sat", "sun":
		fmt.Println("majay")
	default:
		fmt.Println("Not day")
	}

	temp := 80

	switch {
	case temp > 80:
		fmt.Println("hot")
	case temp < 80:
		fmt.Println("cold")
	case temp == 80:
		fmt.Println("hot-cold")
	default:
		fmt.Println("ja ghari")
	}

}
