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
		// we can print out single line commands like ls
		// but commands with args need extra work
		// e.g. echo text or tr "a-z" "A-Z" < my_text.txt
		// mv test.py new_test.py
		if err != nil {
			log.Fatal(err)
		}
		var out strings.Builder
		cmd.Stdout = &out
		cmd.Run()
		fmt.Printf("%v", out.String())
	}

}
