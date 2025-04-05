package main

import (
	"bufio"
	"fmt"
	"os"
)

var input string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your full name: ")

	for true {
		scanner.Scan()
		text := scanner.Text()
		input = text
		if input == "exit" {
			break
		}
		fmt.Println("your input: ", text)
	}

}
