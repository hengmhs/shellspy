package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your full name: ")
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(text)
}
