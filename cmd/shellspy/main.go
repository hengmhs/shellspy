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

	for true {
		fmt.Println("input: ")
		scanner.Scan()
		text := scanner.Text()
		input = text
		if input == "exit" {
			break
		}
		fmt.Println("your input: ", text)
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

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os/exec"
// 	"strings"
// )

// func main() {
// 	cmd := exec.Command("ls")
// 	cmd.Stdin = strings.NewReader("some input")
// 	var out strings.Builder
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("in all caps: %q\n", out.String())
// }
