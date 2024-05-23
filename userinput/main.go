package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter your name :-")
	var name string
	fmt.Scan(&name)

	fmt.Println("Your name is ", name)

	reader := bufio.NewReader(os.Stdin)
	name1, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s", name1)
}
