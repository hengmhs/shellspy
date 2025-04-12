package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"shellspy"
	"strings"
)

var input string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Recording session to 'shellspy.txt'")
	for true {
		scanner.Scan()
		text := scanner.Text()
		input = text
		if input == "exit" {
			break
		}
		cmd, err := shellspy.CommandFromString(input)
		if err != nil {
			log.Fatal(err)
		}
		var out strings.Builder
		cmd.Stdout = &out
		cmd.Run()
		fmt.Printf("%v", out.String())
	}

}
